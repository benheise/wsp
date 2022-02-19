package proxy

import (
	"io"

	"github.com/gowsp/wsp/pkg/msg"
)

type Proxy interface {
	Wan() *Wan

	Routing() *Routing

	NewConn(*msg.WspMessage) error

	io.Closer
}
