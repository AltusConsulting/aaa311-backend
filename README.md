# AAA311-Backend

This backend project is used with the [Smart cities](https://github.com/AltusConsulting/smart-cities) application to authenticate and authorize users.


## Getting Started

To start using this project first you must install [Golang](https://golang.org/dl/). 


### Prerequisites

The following dependencies are used in this project:

* [Gin-Gonic](https://github.com/gin-gonic/gin) - HTTP Web Framework 
* [Viper](https://github.com/spf13/viper) - Library used for giving a configuration solution
* [golang-Scribble](https://github.com/nanobox-io/golang-scribble) - Small embedded database used when creating users.


Some of the endpoints of this backend make use of a smtp server to send emails to the users created in its database. In the `config.toml` file you should fill out the email address that will send the emails and the address of your smtp server:

```toml
[jwt]
secret = "MyS3cr3tW0rd"

[server]
port = 9004

[database]
dir = "./database"

[mail]
sender = "<put your email address here>"
smtp = "<put your smtp email address here"
```

There are other options that you can change in this configuration file such as your secret word for the creation and reading of jwt tokens, the port in which the backend will run and the folder where the users will be stored of the embedded database.


### Installing

Once Golang is installed clone this repository and run the following command at your project root
```
go get
```
This will download all the dependencies required. 

## Deployment

Once you have installed all dependencies, go to the project's root and execute the following command in your terminal:

```
go run main.go
```

You could use [Postman](https://www.getpostman.com/postman) for testing all endpoints of this backend.

## Contributing

...

## Versioning

This is the initial version. We use [SemVer](http://semver.org/) for versioning.

## Authors

* [Altus Consulting Software Team](https://github.com/AltusConsulting)

## License

...
