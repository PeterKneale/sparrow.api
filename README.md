# Mango
Mango Project

## Generation options
```
goagen bootstrap -d github.com/simplicate/sparrow.api/design -o temp
```
## Run locally
```
export DB_HOST=localhost
export DB_DATABASE=sparrow
export DB_USER=postgres
export DB_PASSWORD=password
export DB_PORT=5432
./sparrow.api.exe

```
## Build docker
```
env GOOS=linux GOARCH=386 go build
docker build -t sparrow_api .
```

## Run DB
```
docker run --name db \
    -e POSTGRES_DB=db \
    -e POSTGRES_USER=username \
    -e POSTGRES_PASSWORD=password \
    -d postgres

docker logs db
docker inspect db | grep IPAddress
```


## Run API
```
docker run --name api \
    -e DB_HOST=172.17.0.2 \
    -e DB_DATABASE=db \
    -e DB_USER=username \
    -e DB_PASSWORD=password \
    -e DB_PORT=5432 \
    -d sparrow_api 

docker logs api
docker inspect api | grep IPAddress
```

## Other Commands

### Setup docker machine environment
```
docker-machine env
eval $("C:\Program Files\Docker Toolbox\docker-machine.exe" env)
```

### Stop and remove all docker instances
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```