# ZAPP
Is a scaffolding generator that will produce basic CRUD project complete with server and client code based on script input. Currently server code is written in Golang Language and client code is written in SPA VueJS Framework.

DISCLAIMER: This is still experimental project!


## PURPOSE
* Accomodate your laziness to start coding
* Bootstrap your skeleton code
* Have fix consistent folder structure
* Learn how to code


## PREREQUISITE
Make sure you already install Golang and NPM in your system. Then `git clone` this project and follow those step below for generate, compile and run it


## WRITE YOUR PROJECT
Notice that you have template script in `app/skrip.yaml` that describe the model you want to generate. Some of this field should be self explanatory and you can experiment it by changing the field or create another entities


## GENERATE THE PROJECT
Open this cloned project with your IDE (Visual Studio Code is recommended) and then run with console
```
cd app/
go run main.go
```
At the first time you call this, it will take some time maybe 5 minutes because Golang dependency is being downloaded
The project will generated in `$GOPATH/src/bitbucket.org/mirzaakhena/experimenfx1`. Under that path, you will see there is 2 directory which is `client/` and `server/`


## WORKING ON DEVELOPMENT MODE
Open that new project with your IDE And you need 2 console

#1 console will run the frontend side

Go to `client/` directory
```
cd client/
```

Download webapp dependency. It will take some time to download web dependency
```
npm install
```
Make sure `node_modules/` is created dependency is success downloaded


Run the client in development mode
```
npm run serve
```

#2 console will run the backend side

Go to `server/app/` directory
```
cd server/app/
```

Run the server application
```
go run main.go
```

Open your browser then access the `http://localhost:4000`

In the beginning, you will see the login page. Click the register button then register your self there. Back to login page, then login with the username and password. The API available in `http://localhost:8081`

By default it use SQLite DB. But you can easily change it into other database that supported by GORM.



## BUILD FOR PRODUCTION

Go to `client/` directory
```
cd client/
```

Build it
```
npm run build
```

Client distribution file will be on `client/dist/`

Go to `server/` directory
```
cd server/
```

Build it
```
./build.sh
```

Server executable file will be on `server/dist/`

Run it by
```
cd server/dist/
./experimenfx1
```

Open browser then access `http://localhost:8081`

## RULES FOR SERVER (BACKEND)

### Repository
Is the source of data. It can be database, file, webservice, or anything.

### Service
Is the place you put the logic of your application

### Controller
Is the place where you publish your application as REST API to be used by external resource


## NEXT IMPROVEMENT
* Error Codes
* SystemUser
* -e configfile
