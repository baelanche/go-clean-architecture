# Update swagger
swag init -g ./app/cmd/main.go

# Run
docker-compose -f ./app/docker/docker-compose.yml up