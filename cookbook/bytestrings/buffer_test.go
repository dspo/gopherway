package bytestrings

import (
	"testing"
	"bytes"
	"fmt"
	"bufio"
)


func TestToString(t *testing.T) {
	rawString := "this is a test stirng"
	r := bytes.NewBufferString(rawString)
	res, _ := ToString(r)
	//检查是否转换成功
	t.Log(res)
	if res != rawString {
		t.Error("convert failed !")
	}

	//检查是否被消耗
	res, _ = ToString(r)
	t.Log(res)
	if res != "" {
		t.Error("consume failed !")
	}
}

func TestBuffer(t *testing.T) {
	rawString := "it's easy to encode unicode into a byte array"
	b := Buffer(rawString) //将原始string转换为*bytes.Buffer
	s, err := ToString(b) //*bytes.Buffer转换为string
	if err != nil {
		t.Error(err)
	}
	fmt.Println(b.String())  //调用*bytes.Buffer.String()方法将得到*bytes.Buffer的string
	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))  //得到一个bytes.Reader，它实现了io.Reader接口
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
		fmt.Print(scanner.Text())
	}
}