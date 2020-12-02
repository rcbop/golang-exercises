package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// TODO
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// >>>>>>>>>

	// body, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", body)

	io.Copy(os.Stdout, resp.Body)
}
