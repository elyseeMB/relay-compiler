.PHONY: dev

dev: ## start environment develop
	@make -j2 frontend backend

frontend:
	pnpm run dev

backend:
	go run ./server.go