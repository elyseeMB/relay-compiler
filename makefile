DOCKER?= docker
DOCKER_COMPOSE= $(DOCKER) compose -f compose.yml $(DOCKER_COMPOSE_FLAGS)


.PHONY: dev
dev: ## start environment develop
	@make -j2 stack-up backend frontend

frontend:
	pnpm run dev


.PHONY: backend
backend:
	go run cmd/relay-compiler/main.go

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen generate


.PHONY: stack-up
stack-up:
	$(DOCKER_COMPOSE) up -d


.PHONY: stack-down
stack-down:
	$(DOCKER_COMPOSE) down


.PHONY: stack-ps
stack-ps:
	$(DOCKER_COMPOSE) ps

.PHONY: psql
psql:
	$(DOCKER_COMPOSE) exec postgres psql -U tp -d tp_database

	
