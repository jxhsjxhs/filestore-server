package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

func main()  {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/file/upload",handler.UploadHandler)
	err:= http.ListenAndServe(":8080",nil)
	if err == nil{
		fmt.Println("err",err)
	}
}