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
	docker compose logs traefik --follow

.PHONY: down
down:
	docker compose down --remove-orphans --volumes

.PHONY: sort
sort:
	@yq -i -P 'sort_keys(..)' docker-compose.yaml
	@yq -i -P '(.services[] | select(.volumes != null).volumes) |= sort' docker-compose.yaml
	@yq -i -P '(.services[] | select(.labels != null).labels) |= sort' docker-compose.yaml
	@yq -i -P 'sort_keys(..)' traefik.yaml 
