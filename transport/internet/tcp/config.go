package tcp

import (
	"github.com/gptlocal/netools/common"
	"github.com/gptlocal/netools/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
