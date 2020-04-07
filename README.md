# How To use
Git clone this project

## GENERATE YOUR PROJECT
Notice that you have `app/skrip.txt` that describe the model you want to generated


## COMPILE THE PROJECT
Open this project with your IDE and then run with console
```
cd app/
go run main.go
```
Then project will generated in `$GOPATH/src/github.com/mirzaakhena/sample`


## WORKING ON DEVELOPMENT MODE
Open that new project with your IDE And you need 2 console

First console to run the frontend side

Go to `client/` directory
```
cd client/
```

Download dependency
```
npm install
```

Run the client
```
npm run serve
```

Second console to run the backend side

Go to `server/app/` directory
```
cd server/app/
```

Run the server application
```
go run main.go
```

Open browser then access `http://localhost:4000`


## BUILD FOR PRODUCTION

Go to `client/` directory
```
cd client/
```

Build it
```
npm run build
```

Then output will be on `dist/`




Open browser then access `http://localhost:8081`
