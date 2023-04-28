package all

import (
	"github.com/gptlocal/netools/main/commands/base"
)

//go:generate go run github.com/gptlocal/netools/common/errors/codegen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		//api.CmdAPI,
		// cmdConvert,
		//tls.CmdTLS,
		cmdUUID,
		cmdX25519,
	)
}
