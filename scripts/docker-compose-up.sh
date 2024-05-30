#/bin/sh

docker compose \
	--project-directory . \
	--file ./deployments/docker-compose/docker-compose.local.yaml \
	up