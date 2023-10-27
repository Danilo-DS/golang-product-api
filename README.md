# Product Api

This project is a simple Product and Category API with standard operations (CRUD). There's nothing too complex about it, but it was created to practice my learning in golang.

For this project I'm using some technologies and good practices such as:

- Relational Database (Mysql)
- Migrations
- Design Pattern (Factory Method)
- DTO (Data Object Transfer)
- Middlewares
- Generics
- Docker

## Struture Package
root
 |__config
 |  |__db
 |__controller
 |__dto
 |__mapper
 |__middlewares
 |__models
 |__repository
 |__routers
 |__services
 |__utils


## Instruction to run application

### Local

1. Install and configure Golang on your computer.
2. Download the project and extract the zip file or fork this project from github and clone it.
3. After downloading or cloning, enter the `root > src` folder.
4. Run this command: `go run ./server.go`

> **OBS:** You need to have the Mysql database installed or use docker to run Mysql via the docker image on your computer..

### Docker

> **A little knowledge of Docker is required.** 

1. Download the project and extract the zip file or fork this project from github and clone it.
2. After downloading or cloning, enter the `root` folder.
3. Run this command `docker build  -t golang-product-api -f ./docker/Dockerfile ./`, this command generates the docker iamge of the application.
4. To run the application, use this command: 
`docker run -d --name product-api-go -p 5000:5000 -e USER_DATABASE=productdb -e DATABASE_HOST=mysqldb --network golang_network golang-product-api`.

> **OBS:** 
>- You need to have docker installed on your computer.
>- For this application to connect with database both they must both be on the same network.
>- In the dockerfile there are some environment variables that can be used in the command mentioned in step 4. 
    Environment avaliables:
    - USER_DATABASE
        defalut value: root
    - PASSWORD_DATABASE
        defalut value: root
    - DATABASE_NAME
        defalut value: productdb
    - DATABASE_HOST
        defalut value: localhost
    - DATABASE_PORT
        defalut value: 3306

