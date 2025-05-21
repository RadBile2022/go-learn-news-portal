package main

import (
	"fmt"
	"github.com/RadBile2022/go-learn-news-portal/client/lib"
)

func main() {
	// Login
	token, err := lib.Login("http://localhost:11112/api/login", "admin@gmail.com", "12345678")
	if err != nil {
		fmt.Println("Login Error:", err)
		return
	}
	fmt.Println("Token:", token)

	// Pilih salah satu: file lokal
	//source := lib.UploadSource{
	//	LocalFiles: map[string]string{
	//		"file1": "./files/hutan.jpg",
	//		"file2": "./files/hutan_grayscale.png",
	//	},
	//}

	////Nanti kalau dari handler tinggal isi UploadedFiles (Gin/Fiber)
	//source := lib.UploadSource{
	//	UploadedFiles: form.File,
	//}

	//var fields []string
	//for k := range uploadSource.UploadedFiles {
	//	fields = append(fields, k)
	//}

	//body, contentType, err := lib.PrepareMultipartUnified(fields, source)
	//body, contentType, err := lib.PrepareMultipartFromHeaders([]multipart.FileHeader{})
	//if err != nil {
	//	fmt.Println("Prepare Multipart Error:", err)
	//	return
	//}
	//
	//if err := lib.Upload(token, body, contentType); err != nil {
	//	fmt.Println("Upload Error:", err)
	//	return
	//}
}
