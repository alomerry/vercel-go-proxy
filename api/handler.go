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

	fmt.Printf(`
u: [` + u.Path + `]
req.URL.RawPath:[` + r.URL.RawPath + `],
req.URL.RawQuery:[` + r.URL.RawQuery + `],
req.URL.Path:[` + r.URL.Path + `],
resp.header:[` + getHeaders(w) + `],
	`)
}

func getHeaders(w http.ResponseWriter) string {
	var result []string
	for k, v := range w.Header() {
		result = append(result, strings.Join(append([]string{k}, v...), ","))
	}

	return strings.Join(result, "\n")
}
