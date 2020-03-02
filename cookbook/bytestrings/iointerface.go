package bytestrings

import (
	"io"
)

type IOReaderImp struct {
	Data []byte
}

func (i *IOReaderImp) Read(p []byte) (n int, err error) {
	lenData := len(i.Data)
	if lenData == 0 {
		return 0, io.EOF
	}

	lenp := len(p)
	if lenData <= lenp {
		n := copy(p, i.Data)
		i.Data = []byte{}
		return n, io.EOF
	}

	p, i.Data = i.Data[:len(p)], i.Data[len(p):]
	return lenp, nil
}
