package transfer

import data "proxy-download/src/proxy"

type Transfer interface {
	Handle(p *data.Proxy)
}

var transfers = make([]Transfer, 10)

func Register(transfer Transfer) {
	transfers = append(transfers, transfer)
}
