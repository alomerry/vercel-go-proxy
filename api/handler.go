package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	u, _ := url.Parse("https://api.openai.com/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)

	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", w)
}

func getHeaders(w http.ResponseWriter) string {
	var result []string
	for k, v := range w.Header() {
		result = append(result, strings.Join(append([]string{k}, v...), ","))
	}

	return strings.Join(result, "\n")
}
