package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"url_shortener/url"
)

var (
	logOn   *bool
	port    *int
	baseUrl string
)

func init() {
	domain := flag.String("d", "localhost", "domain")
	port = flag.Int("p", 8888, "port")
	logOn = flag.Bool("l", true, "log on/off")

	flag.Parse()

	baseUrl = fmt.Sprintf("http://%s:%d", *domain, *port)
}

type Headers map[string]string

type Redirector struct {
	stats chan string
}

func (r *Redirector) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	searchUrlToExecute(w, req, func(url *url.Url) {
		http.Redirect(w, req, url.Destination, http.StatusMovedPermanently)
	})
}

func Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respWith(w, http.StatusMethodNotAllowed, Headers{"Allow": "POST"})
		return
	}

	url, new, err := url.SearchOrCreateNewUrl(extractUrl(r))

	if err != nil {
		respWith(w, http.StatusBadRequest, nil)
		return
	}

	var status int
	if new {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortUrl := fmt.Sprintf("%s/r/%s", baseUrl, url.Id)

	respWith(w, status, Headers{
		"Location": shortUrl,
		"Link":     fmt.Sprintf("<%s/api/stats/%s; rel=\"stats\"", baseUrl, url.Id),
	})
}

func Shower(w http.ResponseWriter, r *http.Request) {
	searchUrlToExecute(w, r, func(url *url.Url) {
		json, err := json.Marshal(url.Stats())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respWithJSON(w, string(json))
	})
}

func searchUrlToExecute(w http.ResponseWriter, r *http.Request, executor func(*url.Url)) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if url := url.Search(id); url != nil {
		executor(url)
	} else {
		http.NotFound(w, r)
	}
}

func respWith(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func respWithJSON(w http.ResponseWriter, resp string) {
	respWith(w, http.StatusOK, Headers{"Content-Type": "application/json"})
	fmt.Fprint(w, resp)
}

func extractUrl(r *http.Request) string {
	rawBody := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(rawBody)
	return string(rawBody)
}

func registerStats(stats <-chan string) {
	for id := range stats {
		url.RegisterClick(id)
		registerLog("Click registered with success to %s.", id)
	}
}

func registerLog(format string, values ...interface{}) {
	if *logOn {
		log.Printf(fmt.Sprintf("%s\n", format), values...)
	}
}

func main() {
	url.SetRepository(url.NewMemoryRepository())

	stats := make(chan string)
	defer close(stats)
	go registerStats(stats)

	http.Handle("/r/", &Redirector{stats})
	http.HandleFunc("/api/short", Shortener)
	http.HandleFunc("/api/stats/", Shower)

	registerLog("Starting server on port %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
