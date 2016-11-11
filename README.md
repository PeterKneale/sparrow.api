# Sparrow
Sparrow API

## Publish Docker Images to Registry

 - Build the `debug` docker image using the GOLANG base image (~ 600MB, good for diagnostics)
    ```
    go build
    docker build -t sparrow/api-debug -f Dockerfile.debug .
    docker tag sparrow/api-debug gcr.io/simplicate-sparrow-project/api-debug
    gcloud docker -- push gcr.io/simplicate-sparrow-project/api-debug
    ```

- Build the `release` docker image using the ALPINE base image. (~ 10MB and note it is compiled differently)
    ```
    CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"'
    docker build -t sparrow/api-release -f Dockerfile.release .
    docker tag sparrow/api-release gcr.io/simplicate-sparrow-project/api-release
    gcloud docker -- push gcr.io/simplicate-sparrow-project/api-release
    ```


### Run
```
docker run --name db \
    -d postgres \
    -p 5432:5432 

docker run --name api \
    --link db:db \
    -d sparrow/api-debug \ 
    -p 80:80 
```

### STOP
- delete local containers and images
    ```
    docker rm -f $(docker ps -a -q)
    docker rmi -f $(docker images -q)
    ```
