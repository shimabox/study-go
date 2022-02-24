.PHONY: build
build:
	docker-compose up -d --build

.PHONY: up
up:
	docker-compose up -d

.PHONY: in
in:
	docker-compose exec go bash

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: destroy
destroy:
	docker-compose down --rmi all --volumes --remove-orphans