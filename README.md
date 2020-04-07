# ZAPP
Is a scaffolding generator that will produce basic CRUD project complete with server and client code based on some input. Server code is written in Golang Language and client code is written in SPA VueJS Framework.

DISCLAIMER: This is still experimental project

## PREREQUISITE

Make sure you already install Golang and NPM in your system. Then just git clone this project

## WRITE YOUR PROJECT
Notice that you have template script in `app/skrip.yaml` that describe the model you want to generate. Some of this field should be self explanatory and you can experiment it by changing the field or create another entities


## GENERATE THE PROJECT
Open this cloned project with your IDE and then run with console
```
cd app/
go run main.go
```
Then project will generated in `$GOPATH/src/bitbucket.org/mirzaakhena/experimenfx1`. Under that path, You will see there is 2 directory which is `client/` and `server/`


## WORKING ON DEVELOPMENT MODE
Open that new project with your IDE And you need 2 console

First console will run the frontend side

Go to `client/` directory
```
cd client/
```

Download webapp dependency
```
npm install
```

This client code currently written in vuejs framework

Run the client in development mode
```
npm run serve
```

Second console will run the backend side

Go to `server/app/` directory
```
cd server/app/
```

This server code currently written in golang code

Run the server application
```
go run main.go
```

Open your browser then access the `http://localhost:4000`

In the beginning, you will see the login page. Click the register button then register your self there. Back to login page, then login with the username and password.

By default it use SQLite DB. But you can easily change it into Other Database that supported by GORM.



## BUILD FOR PRODUCTION

Go to `client/` directory
```
cd client/
```

Build it
```
npm run build
```

Then client file output will be on `client/dist/`

Go to `server/` directory
```
cd server/
```

Build it
```
./build.sh
```

Then server executable output will be on `server/dist/`

Run it by
```
cd server/dist/
./experimenfx1
```

Open browser then access `http://localhost:8081`
