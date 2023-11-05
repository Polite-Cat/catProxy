package main

import (
	"fmt"
	"github.com/networm6/catProxy/confbox"
	"github.com/networm6/catProxy/data"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func proxyPort(list data.PointList) {
	points := list.Points
	port := list.Port
	proxyMap := make(map[string]*httputil.ReverseProxy)
	for _, v := range points {
		v := string(v)
		parse, _ := url.Parse(v)
		proxyMap[v] = httputil.NewSingleHostReverseProxy(parse)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			srcUrl = data.SrcUrl(r.Host)
			dstUrl = string(points[srcUrl])
		)
		r.Host = dstUrl
		proxyMap[dstUrl].ServeHTTP(w, r)
	}))
	log.Fatalf("http reverse proxy %d %v", port, err)
}

func main() {
	var list data.PortList
	err := confbox.Load("I:\\ReverseProxy\\test.yaml", &list)
	if err != nil {
		log.Fatalln(err)
	}
	for _, pointList := range list.PointLists {
		go proxyPort(pointList)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
