package main

import (
	"net/http"

	// Import the internal/translations package, so that its init()
	// function is called
	_ "github.com/favtuts/golang-i18n-bookstore/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Extract the locale from the URL path.
	locale := r.URL.Query().Get(":locale")

	// Declare variable to hold the target language tag.
	var lang language.Tag

	// Use language.MustParse() to assign the appropriate language tag
	// for the locale
	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	case "fr-ch":
		lang = language.MustParse("fr-CH")
	default:
		http.NotFound(w, r)
		return
	}

	// Define a variable to hold the number of books. In a real application
	// this would probably be retrieved by making a database query or
	// something similar.
	var totalBookCount = 1

	// Initialize a message.Printer which uses the target lanuage.
	p := message.NewPrinter(lang)
	// Print the welcome message translated into the target language.
	p.Fprintf(w, "Welcome!\n")

	// Use the Fprintf() function to include the new message in the HTTP
	// response, with the book count as in interpolated integer value.
	p.Fprintf(w, "%d books available\n", totalBookCount)
}
