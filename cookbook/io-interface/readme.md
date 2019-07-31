# I/O
## I/O操作核心接口: io.Reader & io.Writer
![](/img/io.Reader++io.Writer+interfaces.svg)

## 使用临时文件
以下是IO操作中使用临时文件的例子
```go
func someFunc() {
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return
	}
	defer os.RemoveAll(t)
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return
	}
	fmt.Println(tf.Name())
}
```