# Marvik Challenge

This is a solution of the Marvik Challenge 

The application is Dockerized, has a docker-compose file to boot the service quickly and a Makefile to start working ASAP.

First you'll have to clone the repo.

```bash
git clone https://github.com/SolBaa/marvik_challenge.git
```

Inside marvik_challenge folder run:

	make init
	make build


`init` will populate the `.env` file needed for injecting environment variables

`build` will create the development image to code inside of it.  
  


If you want to populate the database with data you have to get inside the database running:

      make database

 and then you have to copy the sql statements found in Docs/data.sql file (in the db.sql file you will find the entire database)

Finally run:

    make up
`up` will run the API, exposing ports specified in the docker-compose file.  

And you are ready to Go!

---


### Database

I'll use Postgres as a Database.

### Documentation

[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/10470329-a5c72c4d-8d69-409f-b923-745224eff2c5?action=collection%2Ffork&collection-url=entityId%3D10470329-a5c72c4d-8d69-409f-b923-745224eff2c5%26entityType%3Dcollection%26workspaceId%3Df55dfa65-4072-4bc6-a31b-7dcd012dc208)

(If the button doesn't work, inside Docs folder you'll find the postman collection in JSON format so you can import it in your postman)

### Create Users
```
POST - localhost:8080/users
```
#### Body
```bash
{
  "name": "Matias",
  "last_name": "Lopez",
  "email": "mati@gmail.com"
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
company

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

``` bash
{
    "get_users": 1,
    "get_user_by_id": 2,
    "delete_user": 1,
    "endpoint_counter": 3
}

```


###  Every Five Minutes you'll find a logger that tells you how many minutes have passed since the service is up.

``` shell
INFO    2021/12/09 11:11:58 5 more minutes have passed since I'm Up, time: 5 minutes
INFO    2021/12/09 11:16:58 5 more minutes have passed since I'm Up, time: 10 minutes
INFO    2021/12/09 11:21:58 5 more minutes have passed since I'm Up, time: 15 minutes
```