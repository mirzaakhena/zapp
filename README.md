# GENERATE YOUR PROJECT
Notice that you have `app/skrip.txt` that describe the model you want to generated


# COMPILE THE PROJECT
Open this project with your IDE and then run with console
```
cd app/
go run main.go
```
Then project will generated in `$GOPATH/src/github.com/mirzaakhena/sample


# WORKING ON GENERATED CODE
Open that new project with your IDE And you need 2 console to run in development mode

First console
```
cd webapp/
npm install
npm run build
npm run serve
```

Second console

```
cd app/
go run main.go
```

# In Development Mode
Open browser then access `http://localhost:4000`

# In Production Mode
Open browser then access `http://localhost:8081`
