package main

import(
	"fmt"
	"net/http"

)

func main()  {
	resp, err := http.Get("http://example.com/")
	fmt.Println(resp, err)
	resp1, err1 := http.Post("http://example.com/upload", "image/jpeg", &buf)
	fmt.Println(resp1, err1)
}