package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	data,err := ioutil.ReadFile("./static/view/index.html")
	if err !=nil{
		fmt.Println("err",err)
		return
	}
	io.WriteString(w,string(data))
	} else if  r.Method =="POST"{

	}
}
