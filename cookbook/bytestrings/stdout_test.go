package bytestrings

import (
	"os"
	"testing"
)

func TestStdout(t *testing.T){
	os.Stdout.Write([]byte("test Stdout print"))
}
