.PHONY: up
up:
	docker compose up --detach --remove-orphans	

.PHONY: stop
stop:
	docker compose stop

.PHONY: restart
restart: stop up

.PHONY: logs
logs:
	docker compose logs --follow

.PHONY: down
down:
	docker compose down --remove-orphans --volumes
