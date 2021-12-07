# Marvik Challenge

<!-- This is a solution of the Marvik Challenge 

The application is Dockerized, has a docker-compose file to boot the service quickly and a Makefile to start working ASAP.

	make build
	make up

And you are ready to Go!
 
`build` will create the development image to code inside of it.  
`up` will run the API, exposing ports specified in the docker-compose file.  

This lets the developer focus on the code, running it inside the container resembling production. -->

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
Documentation of the endpoints will be done using OpenAPI spec in Swagger format. 
