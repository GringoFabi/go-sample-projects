# go-sample-projects
A collection of some small sample projects written in Go

## database-app
This module contains a small go app which communicates with a postgres database.
The database itself is run inside a docker-container (see here: [docker-compose.yml](database-app/docker-compose.yml)).

## go-docker
The ``go-docker`` module contains a small go server built with the [gin framework](https://github.com/gin-gonic/gin).
The server can be deployed through the given [Dockerfile](go-docker/Dockerfile). 
Alternatively, the docker-image and the respective containers can be built through the [docker-compose.yml](go-docker/docker-compose.yml).
For more information you may visit this [tutorial](https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e)!

Use the commands listed below to build the image and run your container:
```commandline
docker build . -t go-docker
docker run -p 3000:3000 go-docker
```

## go-graphql
The module contains a very small sample app which introduces and showcases basic concepts of GraphQL and their usage in Golang. 

## rest-server
This module contains a sketch of a rest-server built with the [gorilla framework](https://github.com/gorilla).
Apart from that, this sample project contains some basic methods for working with self-declared data-classes (e.g. an [``Article`` class](rest-server/data.go)).

## websocket
The ``websocket`` module contains a sketch of a websocket built with the [gorilla framework](https://github.com/gorilla).
The running server can be accessed through the given [``index.html``](websocket/index.html) file.
