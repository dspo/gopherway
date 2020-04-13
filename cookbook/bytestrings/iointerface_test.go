package bytestrings

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestIOReaderImp_Read(t *testing.T) {
	var ioReader = &IOReaderImp{Data:[]byte(strings.Repeat("i am testing ", 100))}
	data, err := ioutil.ReadAll(ioReader)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}
