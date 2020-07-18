package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main()  {
	//提供静态文件中的css以及js访问
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadsucHandler)
	err:= http.ListenAndServe(":8080",nil)
	if err == nil{
		fmt.Println("err",err)
	}
}