# Sparrow
Sparrow Project

## Executing via docker-compose

### Setup environment and execute UP
export DB_DATABASE=sparrowdatabase
export DB_USER=sparrowuser
export DB_PASSWORD=sparrowpassword
export DB_PORT=5432

docker-compose up

## Executing LOCALLY

### Run API against local DB
```
export DB_HOST=localhost
export DB_DATABASE=sparrow
export DB_USER=postgres
export DB_PASSWORD=password
export DB_PORT=5432
./sparrow.api.exe
```

## Running in Docker

### Setup docker machine environment and update the shell to use it
```
docker-machine create --driver virtualbox dev
eval $("C:\Program Files\Docker Toolbox\docker-machine.exe" env dev) 
docker-machine ip dev

```

### Build for linux
```
env GOOS=linux GOARCH=386 go build
docker build -t api .
```

### Run DB in docker
```
docker run --name db \
    -e POSTGRES_DB=db \
    -e POSTGRES_USER=username \
    -e POSTGRES_PASSWORD=password \
    -p 5432:5432 \
    -d postgres

docker logs db
docker inspect db | grep IPAddress
```

### Run API in docker
```
docker run --name api \
    -e DB_HOST=172.17.0.2 \
    -e DB_DATABASE=db \
    -e DB_USER=username \
    -e DB_PASSWORD=password \
    -e DB_PORT=5432 \
    -p 8080:8080 \
    -d api 

docker logs api
docker inspect api | grep IPAddress
```

## TEST
```
docker-machine ip dev
curl -v -i -H "Accept: application/json" http://192.168.99.102:8080/users
curl -v -i -H "Accept: application/json" http://192.168.99.102:8080/swagger.json
```

## Other Commands

### Generation of api and models
```
./generate.sh
```

### Stop and remove all docker instances
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```
