.PHONY: up down prune log

up:
	docker compose watch

down:
	docker compose down

prune:
	docker image prune -f

log:
	docker compose logs -f -t