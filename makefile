
DB_HOST=localhost
DB_PORT=5432
DB_USER=tp
DB_PASSWORD=password
DB_NAME=tp_database
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable


PATH_MIGRATION=./pkg/coredata/migrations

DOCKER?= docker
DOCKER_COMPOSE= $(DOCKER) compose -f compose.yml $(DOCKER_COMPOSE_FLAGS)


.PHONY: dev
dev: ## start environment develop
	@make -j2 stack-up frontend backend

frontend:
	pnpm run dev


.PHONY: backend
backend:
	air

.PHONY: clean
clean:
	go clean -cache
	backend


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




.PHONY: migration migrate-up migrate-down migrate-version migrate-force

# Créer une nouvelle migration
# Usage: make migration NAME=create_users_table
migration:
	@if [ "$(NAME)" = "" ]; then \
		echo "❌ Erreur: Spécifiez un nom avec NAME=..."; \
		echo "💡 Usage: make migration NAME=create_users_table"; \
		exit 1; \
	fi
	@echo "📁 Création de la migration: $(NAME)"
	migrate create -ext sql -dir $(PATH_MIGRATION) -seq $(NAME)
	@echo "✅ Migration créée dans $(PATH_MIGRATION)"

# Appliquer toutes les migrations
migrate-up:
	@echo "🚀 Application des migrations..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) up
	@echo "✅ Migrations appliquées"

# Revenir en arrière d'une migration
migrate-down:
	@echo "⬇️  Rollback d'une migration..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) down 1
	@echo "✅ Rollback effectué"

# Voir la version actuelle
migrate-version:
	@echo "📊 Version actuelle des migrations:"
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) version

# Forcer une version (en cas de problème)
# Usage: make migrate-force VERSION=1
migrate-force:
	@if [ "$(VERSION)" = "" ]; then \
		echo "❌ Erreur: Spécifiez une version avec VERSION=..."; \
		echo "💡 Usage: make migrate-force VERSION=1"; \
		exit 1; \
	fi
	@echo "🔧 Force de la version $(VERSION)..."
	migrate -database "$(DB_URL)" -path $(PATH_MIGRATION) force $(VERSION)
	@echo "✅ Version forcée à $(VERSION)"

# Créer le dossier migrations s'il n'existe pas
init-migrations:
	@echo "📁 Création du dossier $(PATH_MIGRATION)..."
	@mkdir -p $(PATH_MIGRATION)
	@echo "✅ Dossier créé"

reset-database:
	migrate -database "$(DB_URL)" -path "$(PATH_MIGRATION)" drop

# Aide
help:
	@echo "🛠️  Commandes disponibles:"
	@echo ""
	@echo "📝 Gestion des migrations:"
	@echo "  make migration NAME=create_users_table  - Créer une nouvelle migration"
	@echo "  make migrate-up                         - Appliquer toutes les migrations"
	@echo "  make migrate-down                       - Revenir en arrière d'une migration"
	@echo "  make migrate-version                    - Voir la version actuelle"
	@echo "  make migrate-force VERSION=1            - Forcer une version"
	@echo "  make init-migrations                    - Créer le dossier migrations"
	@echo ""
	@echo "🔧 Configuration:"
	@echo "  DB_URL: $(DB_URL)"
	@echo "  PATH_MIGRATION: $(PATH_MIGRATION)"

	
