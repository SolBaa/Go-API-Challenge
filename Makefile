default:
	@echo "###################-Mikado Backend Service MakeFile-##########################"
	@echo "- build: Build the service"
	@echo "- devshell: Run the container and open a sh shell inside it"
	@echo "- up: Run the container and executes the service inside"
	@echo "- t: Performs the Unit Testing . Run this inside the container"
	@echo "- down: Take down the service container and all the dependent containers"
	@echo "- migrateup: runs database migrations progressively"
	@echo "- migratedown: destroys database migrations progressively"
	@echo "###################-Mikado Backend Service MakeFile-##########################"

build:
	@docker-compose build --no-cache

devshell:
	@docker-compose run --rm --service-ports api sh

up:
	@docker-compose run --rm --service-ports api

t:
	@go test -cover ./...

down:
	@docker-compose down --remove-orphans
