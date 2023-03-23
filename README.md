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

# working with gotext

Firt need to setup `$GOBIN` environment before running `go install`. Reference: https://stackoverflow.com/questions/25216765/gobin-not-set-cannot-run-go-install
```
$ which go
/usr/local/go/bin/go

$ sudo nano ~/.profile

export GOPATH=$HOME/Projects/go
export GOROOT=/usr/local/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH:$GOBIN


$ source ~/.profile
$ echo $GOBIN
/home/tvt/Projects/go/bin
```

Use `go install` to install the `gotext` executable on your machine:
```
$ go install golang.org/x/text/cmd/gotext@latest
```

All being well, the tool should be installed to your `$GOBIN` directory on your system path and you can run it like so:
```
$ which gotext
/home/tvt/Projects/go/bin/gotext
$ gotext


gotext is a tool for managing text in Go source code.

Usage:

        gotext command [arguments]

The commands are:

        update      merge translations and generate catalog
        extract     extracts strings to be translated from code
        rewrite     rewrites fmt functions to use a message Printer
        generate    generates code to insert translated messages

Use "gotext help [command]" for more information about a command.

Additional help topics:


Use "gotext help [topic]" for more information about that topic.
```


Create that new directory and a `translations.go` file like so:
```
$ mkdir -p internal/translations
$ touch internal/translations/translations.go
```

At this point, your project structure should look like this:
```
$ tree

.
├── README.md
├── cmd
│   └── www
│       ├── handlers.go
│       └── main.go
├── go.mod
├── go.sum
└── internal
    └── translations
        └── translations.go

4 directories, 6 files
```

The content of translactions.go as follow
```go
package translations

//go:generate gotext -srclang=en-GB update -out=catalog.go -lang=en-GB,de-DE,fr-CH github.com/favtuts/golang-i18n-bookstore/cmd/www
```

When we execute this `go generate` command, `gotext` will walk the code for the `cmd/www` application and look for all calls to a `message.Printer`. It then extracts the relevant message strings and outputs them to some JSON files for translation.


OK, let's put this into action and call `go generate` on our `translations.go` file. In turn, this will execute the `gotext` command that we included at the top of that file.

```
$ go generate ./internal/translations/translations.go
de-DE: Missing entry for "Welcome!".
fr-CH: Missing entry for "Welcome!".
```

If you take a look at the directory structure for your project, it should now look like this:

```
$ tree
.
├── README.md
├── cmd
│   └── www
│       ├── handlers.go
│       └── main.go
├── go.mod
├── go.sum
└── internal
    └── translations
        ├── catalog.go
        ├── locales
        │   ├── de-DE
        │   │   └── out.gotext.json
        │   ├── en-GB
        │   │   └── out.gotext.json
        │   └── fr-CH
        │       └── out.gotext.json
        └── translations.go

8 directories, 10 files
```

We can see that the `go generate` command has automatically generated an `internal/translations/catalog.go` file for us (which we'll look at in a minute), and a locales folder containing `out.gotext.json` files for each of our target languages.

# translating the Messages

The workflow for adding a translation goes like this:

1. You generate the `out.gotext.json` files containing the messages which need to be translated (which we've just done).
2. You send these files to a translator, who edits the JSON to include the necessary translations. They then send the updated files back to you.
3. You then save these updated files with the name `messages.gotext.json` in the folder for the appropriate language.

For demonstration purposes, let's quickly simulate this workflow by copying the `out.gotext.json` files to `messages.gotext.json` files, and updating them to include the translated messages like so:

```
$ cp internal/translations/locales/de-DE/out.gotext.json internal/translations/locales/de-DE/messages.gotext.json
$ cp internal/translations/locales/fr-CH/out.gotext.json internal/translations/locales/fr-CH/messages.gotext.json
```

Let's make translaction for messages file, then run our `go generate` command again. This time, it should execute without any warning messages about missing translations.
```
$ go generate ./internal/translations/translations.go
```

The final step in getting this working is to import the `internal/translations` package in our `cmd/www/handlers.go` file. This will ensure that the `init()` function in `internal/translations/catalog.go` is called, and the default message catalog is updated to be the one containing our translations. Because we won't actually be referencing anything in the `internal/translations` package directly, we'll need to alias the import path to the blank identifer `_` to prevent the Go compiler from complaining.

And then run the application:
```
$ go run ./cmd/www/
```

Alright, let's try this out! When your restart the application and try making some requests, you should now see the `"Welcome!"` message translated into the appropriate language.
```
$ curl localhost:4018/en-gb
Welcome!

$ curl localhost:4018/de-de
Willkommen!

$ curl localhost:4018/fr-ch
Bienvenue !
```


# using variables in translations

To demonstrate, we'll update the HTTP response from our `handleHome()` function to include a `"{N} books available"` line, where `{N}` is an integer containing the number of books in our imaginary bookstore.

Then use `go generate` to output some new `out.gotext.json` files. You should see warning messages for the new missing translations like so:
```
$ go generate ./internal/translations/translations.go
de-DE: Missing entry for "{TotalBookCount} books available".
fr-CH: Missing entry for "{TotalBookCount} books available".
```

There is now an entry for our new message. We can see that this has the form `"{TotalBookCount} books available"`, with the (capitalized) variable name from our Go code being used as the placeholder parameter. You should keep this in mind when writing your code, and try to use sensible and descriptive variable names that will make sense to your translators. The `placeholders` array also provides additional information about each placeholder value, the most useful part probably being the `type` value (which in this case tells the translator that the `TotalBookCount` value is an integer).


So the next step is to send these new `out.gotext.json` files off to a translator for translation. Again, we'll simulate that here by copying them to `messages.gotext.json` files and adding the translations like so:
```
$ cp internal/translations/locales/de-DE/out.gotext.json internal/translations/locales/de-DE/messages.gotext.json
$ cp internal/translations/locales/fr-CH/out.gotext.json internal/translations/locales/fr-CH/messages.gotext.json
```

Then run go generate to update our message catalog. This should run without any warnings.
```
$ go generate ./internal/translations/translations.go
```

When you restart the `cmd/www` application and make some HTTP requests again, you should now see the new translated messages like so:
```
$ curl localhost:4018/en-gb
Welcome!
1,252,794 books available

$ curl localhost:4018/de-de
Willkommen!
1.252.794 Bücher erhältlich

$ curl localhost:4018/fr-ch
Bienvenue !
1 252 794 livres disponibles
```

As we'll as the translations being applied by our `message.Printer`, it's also smart enough to output the interpolated integer value with the correct number formatting for each language. We can see here that our `en-GB` locale uses the "," character as a thousands separator, whereas `de-DE` uses "." and `fr-CH` uses the whitespace " ". A similar thing is done for decimal separators too.