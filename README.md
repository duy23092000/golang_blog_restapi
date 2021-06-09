# Go Todo REST API Example
A RESTful API example for simple application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

```bash
# Build and Run
cd tutorial-rest
go get
go build
./tutorial-rest

# API Endpoint : http://127.0.0.1:3000
```

## Structure
```
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── account.go     // APIs for Account model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

#### /projects
* `GET` : Get all projects
* `POST` : Create a new project


