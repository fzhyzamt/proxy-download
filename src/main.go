package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	. "proxy-download/src/config"
	. "proxy-download/src/proxy"
	. "proxy-download/src/transfer"
)

var (
	config Config
)

func index(w http.ResponseWriter, r *http.Request) {
	proxy := &Proxy{w, r, config}
	proxy.Init()

	if proxy.HasHref() {
		var transfer Transfer = DirectTransfer{}
		transfer.Handle(proxy)
	} else {
		proxy.ShowIndexPage()
	}
}

func loadResource(config *Config) {
	path := fmt.Sprintf("../asset/%s", "index.html")
	data, err := Asset(path)
	if err != nil {
		panic("load resources fail")
	}
	config.IndexHtml = string(data)
}

func startServer() {
	addr := fmt.Sprintf(":%d", config.Port)

	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(404) })

	log.Printf("start listen and serve on %d", config.Port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	var (
		port = flag.Int("port", 8081, "http listen port")
		buff = flag.Int("buffer", 8, "download buffer (KB)")
	)
	flag.Parse()

	config = Config{
		BufferSize: 1024 * *buff,
		Port:       *port,
		ExcludeHeaderKey: map[string]bool{
			"Date":   true,
			"Server": true,
			//"Content-Encoding":  true,
			//"Connection":        true,
			//"Transfer-Encoding": true,
		},
	}
	loadResource(&config)

	startServer()
}
