package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"example/richard/sovtech/pkg/models"
	"example/richard/sovtech/pkg/repositories"
)

const (
	baseURI = "https://swapi.dev"
)

var (
	peopleBaseURI  = fmt.Sprintf("%s/api/people/", baseURI)
	planetsBaseURI = fmt.Sprintf("%s/api/planets/", baseURI)
)

type PeopleRepository struct {
	client    *http.Client
	planetMap map[string]string
}

func NewPeopleRepository(client *http.Client) *PeopleRepository {
	pr := &PeopleRepository{client: client}
	pr.getSwapiPlanetMap(context.Background())
	return pr
}

type swapiPeopleResponse struct {
	Next    *string          `json:"next"`
	Results []*models.Person `json:"results"`
}

type swapiPlanetResponse struct {
	Next    *string       `json:"next"`
	Results []swapiPlanet `json:"results"`
}

type swapiPlanet struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (pr *PeopleRepository) getSwapiPlanetMap(ctx context.Context) {
	pr.planetMap = make(map[string]string)
	hasMore := true

	fullURI := fmt.Sprintf("%s", planetsBaseURI)
	for hasMore {
		hasMore = false
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURI, nil)
		if err != nil {
			continue
		}
		resp, err := pr.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
		}
		decoder := json.NewDecoder(resp.Body)
		swapiResponseBody := new(swapiPlanetResponse)
		err = decoder.Decode(swapiResponseBody)
		if err != nil {
			continue
		}
		if swapiResponseBody.Results != nil {
			for _, planet := range swapiResponseBody.Results {
				pr.planetMap[planet.URL] = planet.Name
			}
		}
		if swapiResponseBody.Next != nil {
			fullURI = *swapiResponseBody.Next
			hasMore = true
		}
	}
}

func (pr PeopleRepository) populatePeoplePlanetNames(result *repositories.PeopleResult) {
	if len(pr.planetMap) > 0 {
		for _, person := range result.People {
			if person == nil {
				continue
			}
			personPtr := person
			if planetName, found := pr.planetMap[personPtr.Homeworld]; found {
				personPtr.Homeworld = planetName
			}
		}
	}
}

func (pr PeopleRepository) getSwapiPeople(ctx context.Context, fullURI string) (repositories.PeopleResult, error) {
	result := repositories.PeopleResult{make([]*models.Person, 0), false}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURI, nil)
	if err != nil {
		return result, err
	}
	resp, err := pr.client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("status is %s", resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	swapiResponseBody := new(swapiPeopleResponse)
	err = decoder.Decode(swapiResponseBody)
	if err != nil {
		return result, err
	}
	if swapiResponseBody.Results != nil {
		result.People = swapiResponseBody.Results
	}
	if swapiResponseBody.Next != nil {
		result.HasMore = true
	}

	pr.populatePeoplePlanetNames(&result)
	return result, nil
}

func (pr PeopleRepository) GetPeople(ctx context.Context, page int) (repositories.PeopleResult, error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	fullURI := fmt.Sprintf("%s?%s", peopleBaseURI, params.Encode())

	return pr.getSwapiPeople(ctx, fullURI)
}

func (pr PeopleRepository) SearchPeople(ctx context.Context, name string, page int) (repositories.PeopleResult, error) {
	params := url.Values{}
	params.Add("page", strconv.Itoa(page))
	params.Add("search", name)
	fullURI := fmt.Sprintf("%s?%s", peopleBaseURI, params.Encode())

	return pr.getSwapiPeople(ctx, fullURI)
}
