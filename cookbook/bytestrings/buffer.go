package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

//bytes.Buffer专门用以处理数据。
//在处理流数据时，bytes.Buffer比[]byte更方便。
//bytes.Buffer实现了io.Reader接口，所以可以使用ioutil下的函数处理它。
//在流应用中，尽量使用buffer而不要使用scanner。
//bufio库下有许多相关工具可以使用。
//除非你的应用足够小、或者你的机器内存足够大，否则优先使用buffer而不要使用array或slice。
func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)
	var b = new(bytes.Buffer) //bytes.Buffer实现了io.Reader接口
	b.Write(rawBytes)
	b = bytes.NewBuffer(rawBytes)
	b = bytes.NewBufferString(rawString)
	return b
}

//一个操作io.Reader的示例，
//ToString读取所有数据，消耗所有数据，
//返回读取到的数据（以string形式）。
func ToString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil{
		return "", err
	}
	return string(b), nil
}
