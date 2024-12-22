.PHONY: up stop restart logs down sort

up:
	docker compose up --detach --remove-orphans	

stop:
	docker compose stop

restart: down up

logs:
	docker compose logs traefik --follow

down:
	docker compose down --remove-orphans --volumes

sort:
	@yq -i -P 'sort_keys(..)' docker-compose.yml
	@yq -i -P '(.services[] | select(.volumes != null).volumes) |= sort' docker-compose.yml
	@yq -i -P '(.services[] | select(.labels != null).labels) |= sort' docker-compose.yml
	@yq -i -P 'sort_keys(..)' traefik/dynamic.yml 
	@yq -i -P 'sort_keys(..)' traefik/traefik.yml
	@yq -i -P 'sort_keys(..)' grafana/dashboards.yml
	@yq -i -P 'sort_keys(..)' grafana/datasources.yml
	@yq -i -P 'sort_keys(..)' plugins/request-id/.traefik.yml
	@yq -i -P 'sort_keys(..)' prometheus/prometheus.yml

