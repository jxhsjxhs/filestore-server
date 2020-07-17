package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		file,head,err := r.FormFile("file")
		if err !=nil{
			fmt.Println("err",err)
			return
		}
		defer file.Close()
		newfile,err := os.Create("/tmp/"+head.Filename)
		if err !=nil{
			fmt.Println("err",err)
			return
		}
		defer newfile.Close()
		_,err = io.Copy(newfile,file)
		if err !=nil{
			fmt.Println("err",err)
			return
		}
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
		
	}
}

func UploadsucHandler(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w,"access upload")
}