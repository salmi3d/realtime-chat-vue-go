# Realtime chat
Simple web chat application written with Vue (frontend) and Go (backend).

## Backend
Chat requires a working Go development environment. The [Getting Started](http://golang.org/doc/install) page describes how to install the development environment.
```
cd <project-folder>/backend
go run main.go
```

### Dockerize backend
```
cd <project-folder>/backend
docker build -t backend .
docker run -it -p 4444:4444 backend
```

## Frontend

### Setup
```
cd <project-folder>
yarn install
```
  
### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

To use the chat example, open [http://localhost:8080/](http://localhost:8080/) in your browser
