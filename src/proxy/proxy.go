package data

import (
	"fmt"
	"net/http"
	. "proxy-download/src/config"
	"strings"
)

type Proxy struct {
	W      http.ResponseWriter
	R      *http.Request
	Config Config
}

func (p *Proxy) Init() {
	if err := p.R.ParseForm(); err != nil {
		p.W.WriteHeader(400)
		_, _ = fmt.Fprintf(p.W, "parse request fail\n%s", err.Error())
		return
	}
}

func (p *Proxy) HasHref() bool {
	href := p.R.Form["href"]
	return href != nil && len(href) == 1
}

func (p *Proxy) GetHref() string {
	return p.R.Form["href"][0]
}

func (p *Proxy) ShowIndexPage() {
	p.W.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(p.W, p.Config.IndexHtml)
}

func (p *Proxy) GetHrefFileName() string {
	href := p.GetHref()
	index := strings.LastIndex(href, "/")
	if index == -1 || index+1 == len(href) {
		return ""
	}
	href = href[index+1:]

	return href
}
