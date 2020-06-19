package transfer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	. "proxy-download/src/proxy"
)

func init() {
	Register(DirectTransfer{})
}

type DirectTransfer struct {
}

func (DirectTransfer) Handle(p *Proxy) {
	log.Printf("start download %s", p.GetHref())

	client := http.DefaultClient
	request, err := http.NewRequest("GET", p.GetHref(), nil)
	if err != nil {
		p.W.WriteHeader(500)
		_, _ = fmt.Fprintf(p.W, "create request fail\n%s", err.Error())
		return
	}

	for key, values := range p.R.Header {
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
		p.W.WriteHeader(500)
		_, _ = fmt.Fprintf(p.W, "get source fail\n%s", err.Error())
		return
	}

	// 复制响应头
	for key, values := range resp.Header {
		if !p.Config.ExcludeHeaderKey[key] {
			p.W.Header().Set(key, values[0])
		}
	}
	// 如果响应头没有指定文件名，尝试指定
	// 否则将导致http://example.org/file.txt这种形式的链接无法被浏览器正确的解析文件名
	if _, ok := p.W.Header()["Content-Disposition"]; !ok {
		filename := p.GetHrefFileName()
		if filename != "" {
			p.W.Header().Set("Content-Disposition", "attachment; filename*=UTF-8''"+url.PathEscape(filename))
		}
	}

	if _, err := io.CopyBuffer(p.W, resp.Body, make([]byte, p.Config.BufferSize)); err != nil {
		log.Println("copy data fail", err)
	}

}
