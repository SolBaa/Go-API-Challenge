# Marvik Challenge

This is a solution of the Marvik Challenge 

The application is Dockerized, has a docker-compose file to boot the service quickly and a Makefile to start working ASAP.

	make init
	make build
	make up


`build` will create the development image to code inside of it.  
`up` will run the API, exposing ports specified in the docker-compose file.  

This lets the developer focus on the code, running it inside the container resembling production.


If you want to populate the database with data you have to get inside the database running:

      make database

 and then you have to copy the sql statements found in the db.sql file

And you are ready to Go!

---

<!-- ## Technical Test -->
Implementar una API en Golang que:

a) se comunique con una base de datos (a elección) que tenga una tabla de usuarios con por lo menos los campos nombre, apellido y mail del usuario

b) reciba un request que tenga filtros de búsqueda sobre la tabla usuarios y devuelva la lista de usuarios que cumplen con el criterio de búsqueda, paginando en el backend

c) reciba un request GET que obtenga el valor de un contador de veces que fue llamado alguno de los endpoints disponibles

d) cada 5 minutos incremente un segundo contador e imprima su valor a un log o consola

### Database

I'll use Postgres as a Database.

### Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/10470329-a5c72c4d-8d69-409f-b923-745224eff2c5?action=collection%2Ffork&collection-url=entityId%3D10470329-a5c72c4d-8d69-409f-b923-745224eff2c5%26entityType%3Dcollection%26workspaceId%3Df55dfa65-4072-4bc6-a31b-7dcd012dc208)

### Create Users
```
POST - localhost:8080/users
```
#### Body
```bash
{
  "name": "Matias",
  "last_name": "Lopez",
  "email": "mati@gmail.com",
  "user_company": [
    {
      "name": "Marvik"
    }
  ]
}

```

### Get All Users
```
GET - localhost:8080/users
```
### Get user By ID
```
GET - localhost:8080/users/:userID
```
### Get Users By Filter
```
GET- localhost:8080/users-search?name=sol
```
#### Request params
```bash
LastName
name
email

```
### Create Company
```
POST - localhost:8080/companies
```
#### Body
```bash
{
  "name": "Marvik"
}

```
### Get All Companies
```
GET - localhost:8080/companies
```
### Get Company By ID
```
GET - localhost:8080/companies/:companyID
```


### Add Company To User
```
PUT - localhost:8080/users/:userID/:companyID
```
### Delete User
```
DELETE - localhost:8080/users/userID
```
### Delete All Companies From User
```
DELETE - localhost:8080/users/:userID/companies
```
### Delete A Particular Company From User
```
DELETE - localhost:8080/users/:userId/companies/:companyID
```
### Get Endpoints Count
```
GET - localhost:8080/endpoint-count
```
