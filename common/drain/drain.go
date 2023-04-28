package drain

import "io"

//go:generate go run github.com/gptlocal/netools/common/errors/codegen

type Drainer interface {
	AcknowledgeReceive(size int)
	Drain(reader io.Reader) error
}
