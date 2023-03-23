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

Install third-party for HTTP framework which supports dynamic values in URL path segments.
We are going to use PAT, but feel free to use alternative like [chi](https://github.com/go-chi/chi) or [gorilla/mux](https://github.com/gorilla/mux)

**Note:** If you're not sure which router to use in your project, you might like to take a look at my [comparison of Go routers](https://www.alexedwards.net/blog/which-go-router-should-i-use) blog post.
```
$ go get github.com/bmizerany/pat

go: added github.com/bmizerany/pat v0.0.0-20210406213842-e4b6760bdd6f
```


Add some codes to those two files, Once that's done, run `go mod tidy` to tidy your `go.mod` file and download any necessary dependencies, and then run the web application.
```
$ go mod tidy
$ go run ./cmd/www/
2023/03/23 14:35:57 starting server on :4018...
```

Make some requests to the application using curl, you should find that the appropriate locale is echoed back to you like so:
```
$ curl localhost:4018/en-gb
The locale is en-gb

$ curl localhost:4018/de-de
The locale is de-de

$ curl localhost:4018/fr-ch
The locale is fr-ch

$ curl localhost:4018/da-DK
404 page not found
```


# extracting and translating text content

In this project we'll use British English (en-GB) as the default 'source' or 'base' language in our application, but we'll want to render a translated version of the welcome message in German and French for the other locales. To do this, we'll need to import the [golang.org/x/text/language](https://pkg.go.dev/golang.org/x/text/language) and [golang.org/x/text/message](https://pkg.go.dev/golang.org/x/text/message) packages


After importing and updating codes in your `cmd/www/handlers.go` file. Then run `go mod tidy` to download the necessary dependencies…
```
$ go mod tidy
go: finding module for package golang.org/x/text/message
go: finding module for package golang.org/x/text/language
go: found golang.org/x/text/language in golang.org/x/text v0.8.0
go: found golang.org/x/text/message in golang.org/x/text v0.8.0
```

And then run the application:
```
$ go run ./cmd/www/
2023/03/23 14:54:56 starting server on :4018...
```

When you make a request to any of the supported URLs, you should now see the (untranslated) `welcome` message like this:
```
$ curl localhost:4018/en-gb
Welcome!

$ curl localhost:4018/de-de
Welcome!

$ curl localhost:4018/fr-ch
Welcome!
```

So in all cases we're seeing the "Welcome!" message in our en-GB source language. That's because we still need to provide Go's message package with the actual translations that we want to use. Without the actual translations, it falls back to displaying the message in the source language.