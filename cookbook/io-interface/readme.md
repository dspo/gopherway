# I/O
## I/O�������Ľӿ�: io.Reader & io.Writer
![](/img/io.Reader++io.Writer+interfaces.svg)

## ʹ����ʱ�ļ�
������IO������ʹ����ʱ�ļ�������
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