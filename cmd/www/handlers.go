package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Extract the locale from the URL path. This line of code is likely to
	// be different for you if you are using an alternative router.
	locale := r.URL.Query().Get(":locale")

	// If the locale match one of our supported values, echo the locale
	// in the response. Otherwise send a 404 Not Found response
	switch locale {
	case "en-gb", "de-de", "fr-ch":
		fmt.Fprintf(w, "The locale is %s\n", locale)
	default:
		http.NotFound(w, r)
	}
}
