package handler

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%+v\n", r)
	u, _ := url.Parse("https://api.openai.com/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Director(r)
	fmt.Printf("%+v\n", r)
}
