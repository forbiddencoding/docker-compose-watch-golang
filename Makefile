.PHONY: up donw prune log

up:
	docker compose up --wait && docker compose alpha watch

down:
	docker compose down

prune:
	docker image prune -f

log:
	docker compose logs -f golang