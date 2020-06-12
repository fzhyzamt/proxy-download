package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	// 下载缓冲区大小 (Byte)
	bufferSize       int
	indexHtml        string
	excludeHeaderKey = map[string]bool{
		"Date":   true,
		"Server": true,
		//"Content-Encoding":  true,
		//"Connection":        true,
		//"Transfer-Encoding": true,
	}
)

func init() {
	path := fmt.Sprintf("../asset/%s", "index.html")
	data, err := Asset(path)
	if err != nil {
		panic("load resources fail")
	}
	indexHtml = string(data)
}

type Proxy struct {
	w http.ResponseWriter
	r *http.Request
}

func (p *Proxy) init() {
	if err := p.r.ParseForm(); err != nil {
		p.w.WriteHeader(400)
		_, _ = fmt.Fprintf(p.w, "parse request fail\n%s", err.Error())
		return
	}
}

func (p *Proxy) hasHref() bool {
	href := p.r.Form["href"]
	return href != nil && len(href) == 1
}

func (p *Proxy) getHref() string {
	return p.r.Form["href"][0]
}

func (p *Proxy) showIndexPage() {
	p.w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(p.w, indexHtml)
}

func (p *Proxy) getHrefFileName() string {
	href := p.getHref()
	index := strings.LastIndex(href, "/")
	if index == -1 || index+1 == len(href) {
		return ""
	}
	href = href[index+1:]

	return href
}

func (p *Proxy) downloadAndReWrite() {
	log.Printf("start download %s", p.getHref())

	client := http.DefaultClient
	request, err := http.NewRequest("GET", p.getHref(), nil)
	if err != nil {
		p.w.WriteHeader(500)
		_, _ = fmt.Fprintf(p.w, "create request fail\n%s", err.Error())
		return
	}

	for key, values := range p.r.Header {
		request.Header.Add(key, values[0])
	}

	resp, err := client.Do(request)
	if resp != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Println("close resp body fail", err)
			}
		}()
	}
	if err != nil {
		p.w.WriteHeader(500)
		_, _ = fmt.Fprintf(p.w, "get source fail\n%s", err.Error())
		return
	}

	// 复制响应头
	for key, values := range resp.Header {
		if !excludeHeaderKey[key] {
			p.w.Header().Set(key, values[0])
		}
	}
	// 如果响应头没有指定文件名，尝试指定
	// 否则将导致http://example.org/file.txt这种形式的链接无法被浏览器正确的解析文件名
	if _, ok := p.w.Header()["Content-Disposition"]; !ok {
		filename := p.getHrefFileName()
		if filename != "" {
			p.w.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''"+url.PathEscape(filename))
		}
	}

	if _, err := io.CopyBuffer(p.w, resp.Body, make([]byte, bufferSize)); err != nil {
		log.Println("copy data fail", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	proxy := &Proxy{w, r}
	proxy.init()

	if proxy.hasHref() {
		proxy.downloadAndReWrite()
	} else {
		proxy.showIndexPage()
	}
}

func main() {
	port := flag.Int("port", 8081, "http listen port")
	buff := flag.Int("buffer", 8, "download buffer (KB)")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	bufferSize = 1024 * *buff

	http.HandleFunc("/", index)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(404) })

	log.Printf("start listen and serve on %d", *port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
