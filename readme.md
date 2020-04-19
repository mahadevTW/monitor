Monitoring App serves both Assets and Api

make sure you have go and node installed on local machine. 

**Build UI**
```shell script
cd web
npm install
npm run build
```
Above command will build assets and will keep it in 'web/build' directory

Go back to Project root directory

```shell script
cd ..
```

**Run Monitoring Api**

```go run main.go```

visit http://localhost:8080/

App is by default monitoring 2 URLS, you can go ahead and add more Service URL;s to monitor.

*this app uses in memory database, so restarting app will lose Endpoints being added*