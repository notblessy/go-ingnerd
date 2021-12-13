# go-Ingnerd

go-Ingnerd is a backend todo-list API based on Golang, Gin and Gorm.

## Requirements

- Golang packages
- DBMate
- Gin-gonic
- Gorm

## Features

- CRUD API for todo-list

## How to Run

### Clone

First clone this repo by run:

```sh
$ git clone git@github.com:notblessy/go-ingnerd.git
```

then go to the project folder

```sh
$ cd go-ingnerd
$ go mod init
```

to execute migration proccess, you need to install DBMate **globally**
```sh
$ brew install dbmate
```

and you need to install gin and gorm

```sh
$ go get github.com/gin-gonic/gin github.com/jinzhu/gorm
```

Don't forget to set the .env from **.env.sample** 

## Usage

| `go <script>` | Description                      |
| -------------- | -------------------------------- |
| `run`         | Run & compile go for development  |
| `build`          | Build & compile go for production |

## Additional Packages

- cast - go get github.com/spf13/cast --> to parse dataType
- godotenv - go get github.com/joho/godotenv --> to automatically load variabels from a .env file
## Author

```
Frederich Blessy
```
