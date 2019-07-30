package main

import (
	"encoding/json"
	"fmt"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-web"
	"github.com/prometheus/common/log"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//负责处理/example/call路由
func exampleCall(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	name := r.Form.Get("name")
	if len(name) == 0 {
		http.Error(w,
			errors.BadRequest("go.micro.api.example", "no content").Error(),
			400,
		)
		return
	}

	b, _ := json.Marshal(map[string]interface{}{
		"message": "exampleCall 收到了你的消息, " + name,
	})
	_, _ = w.Write(b)
}

//负责处理/example/foo/bar路由
func exampleFooBar(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w,
			errors.BadRequest("go.micro.api.example", "require post").Error(),
			400)
		return
	}

	if len(r.Header.Get("Content-Type")) == 0 {
		http.Error(w,
			errors.BadRequest("go.micro.api.example", "need content-type").Error(),
			400)
		return
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	data := struct {
		Name string `json:"name"`
	}{}
	_ = json.Unmarshal(bodyBytes, data)
	b, _ := json.Marshal(map[string]interface{}{
		"message": "exampleFooBar收到了你的消息，" + data.Name,
	})

	_, _ = w.Write(b)
}

//负责处理/example/foo/upload上传文件路由
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		text := `<html>
			<head>
				<title>Upload file</title>
			</head>
			<body>
				<form enctype="multipart/form-data" action="http://127.0.0.1:8080/example/foo/upload" method="post">
					<input type="file" name="uploadfile"/>
					<br/>保存目录：<input type="text" name="path"/> 如/Users/me/Downloads/test/<br/>
					<input type="submit" name="上传" value="upload"/>
				</form>
			</body>
			</html>
			`
		t, _ := template.New("foo").Parse(text)
		_ = t.Execute(w, nil)
		return
	}

	_ = r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, _ = fmt.Fprintf(w, "%v", handler.Header)

	path := r.PostForm.Get("path")
	f, err := os.OpenFile(path+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	//为了方便，此处就直接使用go-web包，因为API需要从注册中心获取服务信息，
	//而go-web包已经有注册服务的能力
	service := web.NewService(
		web.Name("go.micro.api.example"),
	)

	service.HandleFunc("/example/call", exampleCall)
	service.HandleFunc("/example/foo/bar", exampleFooBar)
	service.HandleFunc("/example/foo/upload", uploadFile)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
}
