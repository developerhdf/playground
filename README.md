## To Run The Server
```
go run server.go
```
## Supported Query Examples
Remember to include the JWT token header in all requests, for example:
```
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU0MDc3OTEsImlzcyI6InVzZXJAc3dhcGl0ZXN0LmNvbSJ9.AIcmyUTFCrfIgti85Uc9r4jgdU2oQx7yh5Y1raD_oYA"
}
```
### Get People by Page
```
query{getPeople(page:2){people{name homeworld}}}
```
### Get Person by Search Term
```
query{searchPeople(name:"owen"){name height mass homeworld}}
```
### Create User
```
mutation {createUser(input:{email:"user@swapitest.com",password:"newSwapiPwd"})}
```
### Login
```
mutation {login(input:{email:"user@swapitest.com", password:"newSwapiPwd"})}
```
### Token Refresh
```
mutation {refreshToken(input:{token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDU0MDc3NzMsImlzcyI6InVzZXJAc3dhcGl0ZXN0LmNvbSJ9.SobxF-8EIGZtm-LL4xCVMxtfgEo2GaVJ9BzihdpQdJs"})}
```
## Unit Test Example
Only a single unit test file was added with very limited tests.
```
cd pkg/repositories/memory/
go test
```
