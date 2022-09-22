
<h1 align="center">
  Wa-Message App
  <br>
  
</h1>

<h4 align="center">Messaging app Rest API Backend, users can send and reply to messages 
<br>
<br>
<img src="readme\app.gif" alt="Running App" width="70%" height=70% align="center">

</h4>
<p align="left">
<h2>
  Content <br></h2>
  • <a href="#key-features">Key Features</a> <br>
  • <a href="#Installing -and-runing-from-github">Installing Using Github</a> <br>
 • <a href="#installing-and-runing-with-docker">Installing Using Docker</a><br>
  • <a href="#end-point">End Point</a><br>
  • <a href="#technologi">Technologi that i use</a><br>
  • <a href="#contact-me">Contact me</a><br>
</p>


## 📱 Key Features

* User can register
* User can login
* User send messages 
* User reply messages

## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/wa-message.git

# Go into the repository
$ cd wa-app

# Install dependencies
$ go get

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allrady create database for this app see local `.env` file.


## ⚙️ Installing and Runing with Docker
if you are using docker or aws/google cloud server you can run this application by creating a container. <br>

```bash
# Pull this latest app from dockerhub 
$ docker pull programmerreski/wa-app

# if you have mysql container you can skip this
$ docker pull mysql

$ docker run --name mysqlku -p 3306:3306 -d -e MYSQL_ROOT_PASSWORD="yourmysqlpassword" mysql 

# create app container
$ docker run --name waapp -p 80:8000 -d --link mysqlku -e SECRET="secr3t" -e SERVERPORT=8000 -e Name="message_app" -e Address=mysqlku -e Port=3306 -e Username="root" -e Password="yourmysqlpassword" programmerreski/wa-app

# Run the app
$ docker logs waapp
```

## 📜 End Point  

Users
| Methode       | EndPoint       | ...
| ------------- | -------------  | -----------
| `POST`        | /users         | register
| `POST`        | /login         | login
| `GET`         | /users         | get my account
| `DELETE`      | /users         | delete my account
| `PUT`         | /users         | update my account

Message
| Methode       | EndPoint       | ...
| ------------- | -------------  | -----------
| `POST`        | /message       | create message
| `DELETE`      | /message/:id  | delete my message



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

