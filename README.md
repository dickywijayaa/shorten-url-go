# shorten-url-go
Hello, welcome to simple go app for shortening URL with go!

## Steps
Here are some steps if you want to try this repository : 
1. git clone https://github.com/dickywijayaa/shorten-url-go.git

2. cd to the folder by execute : `cd shorten-url-go`

3. install all dependencies by running `go mod download`

4. make sure u also have postgres in local. if you didn't already have postgres, here are the steps you can follow : 
``` https://gist.github.com/ibraheem4/ce5ccd3e4d7a65589ce84f2a3b7c23a3 ```

5. after successfully install postgres in local, copy .env.example and rename it to .env

6. fill the value at env with config as your settings

7. dont forget to run the migrations in folder database/migrations to create the tables.

8. after that, make sure to run : `swag init` so the docs will be generated

9. finally run : `go run main.go` and visit localhost to make sure the server is running! :)

## Contact
dw_authorized@yahoo.co.id