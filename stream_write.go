package openwechat

import (
	"fmt"
)

type Response interface {
	Write(p []byte) (n int, err error)
	Flush() error
}

type StreamWrite struct {
	W Response
}

func (s *StreamWrite) Write(str string) {
	s.W.Write([]byte(fmt.Sprintf("%s\n", str)))
	s.W.Flush()
}
