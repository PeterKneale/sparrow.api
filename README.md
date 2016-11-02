# Sparrow
Sparrow Project

## Generation of api and models
```
./generate.sh
```

## RUNNING LOCALLY
```
./sparrow.api.exe
```


## HOSTING IN DOCKER
### Setup docker machine environment and update the shell to use it
```
docker-machine create --driver virtualbox dev
eval $("C:\Program Files\Docker Toolbox\docker-machine.exe" env dev) 
docker-machine ip dev
```

### Build the API Dockerfile
```
go build
docker build -t api -f Dockerfile.debug
```

### Run DB
```
docker run --name db \
    -d postgres \
    -p 5432:5432 
docker run --name api \
    --link db:db \
    -d api \ 
    -p 80:80 
docker logs api
docker inspect db | grep IPAddress
```

### Stop and remove all docker instances
```
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```
