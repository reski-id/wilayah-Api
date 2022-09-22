

<h1 align="center">
  Wilayah Rest API App
  <br>
  
</h1>

<h4 align="center"> 
<br>
<br>

</h4>
<p align="left">
<h2>
  Content <br></h2>
  • <a href="#Features">Key Features</a> <br>
  • <a href="#github">Installing Using Github</a> <br>
 • <a href="#docker">Installing Using Docker</a><br>
  • <a href="#end-point">End Point</a><br>
  • <a href="#iopenApi">Swagger Open API</a><br>
  • <a href="#technologi">Technologi that i use</a><br>
  • <a href="#contact">Contact me</a><br>
</p>


## 📱 Features

* register
* login
* Provinsi


## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/wilayah-Api.git

# Go into the repository
$ cd wilayah-Api

# Install dependencies
$ go get

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allready create database for this app see local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/wilayah-api

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name wilapi -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="wilayah-api" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/wilayah-api

# Run the app
$ docker logs wilapi
```

## 📜 End Point  

Users
| Methode       | End Point      | used for            | Using JWT
| ------------- | -------------  | -----------         | -----------
| `POST`        | /users         | register            | NO
| `POST`        | /login         | login               | NO 
| `GET`         | /users         | get my account      | YES
| `DELETE`      | /users         | delete my account   | YES
| `PUT`         | /users         | update my account   | YES

Provinsi
| Methode       | End Point       | used for                | Using JWT
| ------------- | -------------   | -----------             | -----------
| `POST`        | /provinsi       | Add New Data Provinsi   | YES
| `GET`         | /provinsi       | get Provinsi            | YES
| `DELETE`      | /provinsi/:id   | Delete Data Provinsi    | YES
| `PUT`         | /provinsi/:id   | Update Data Provinsi    | YES



## 🛠️ Technologi

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Echo Framework](https://echo.labstack.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Mockery Unit Testing](https://github.com/vektra/mockery)
- [Git Repository](https://github.com/reski-id)
- [Trunk Base Development](https://trunkbaseddevelopment.com/)
- [JSON Web Tokens (JWT)](https://jwt.io/)

## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>
