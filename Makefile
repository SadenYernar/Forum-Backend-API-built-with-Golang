run:
    go run cmd/main.go
docker:
    docker build -t forum .
    docker image prune -f 
    docker container prune -f
    docker run -p 8081:8081 --name forum forum