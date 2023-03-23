# golang-i18n-bookstore
Demo for I18n in Go: Managing Translations

Reference
* [I18n in Go: Managing Translations](https://www.alexedwards.net/blog/i18n-managing-translations)


# scaffolding a web application

Setup a new project directory
```
$ mkdir golang-i18n-bookstore
$ cd golang-i18n-bookstore
$ go mod init github.com/favtuts/golang-i18n-bookstore

go: creating new go.mod: module github.com/favtuts/golang-i18n-bookstore
```

Create two file `main.go` and `handlers.go`
```
$ mkdir -p cmd/www
$ touch cmd/www/main.go  cmd/www/handlers.go
```

Check project directory
```
$ sudo apt-get install tree
$ tree
.
├── README.md
├── cmd
│   └── www
│       ├── handlers.go
│       └── main.go
└── go.mod

2 directories, 4 files
```


