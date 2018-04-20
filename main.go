package main

import (
	"fmt"
	"knx/api"
	"knx/db"
	"knx/gui"
	"log"
	"net/http"
	"strings"
)

type Subdomains map[string]http.Handler

func (subdomains Subdomains) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	domainParts := strings.Split(r.Host, ".")

	subdomain := domainParts[0]

	// if no subdomain, redirect to subdomain "api"
	if len(domainParts) < 2 { /*TODO: 3 for real address: www.knx.ru, 2 for test: www.localhost*/
		subdomain = "api"
	}

	if mux := subdomains[subdomain]; mux != nil {
		// Let the appropriate mux serve the request
		mux.ServeHTTP(w, r)
	} else {
		// Handle 404
		fmt.Fprintf(w, "Domain '%s' not served!", subdomain)
		//http.Error(w, "Not found", 404)
	}
}

func main() {
	defer db.DB.Close()

	// Init web-server
	APIMux := api.InitAPIMux()
	GUIMux := gui.InitGUIMux()

	subdomains := make(Subdomains)
	subdomains["api"] = APIMux
	subdomains["www"] = GUIMux

	log.Fatal(http.ListenAndServe(":8080", subdomains))
}
