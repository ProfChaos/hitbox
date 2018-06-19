package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(b))
	})
	http.ListenAndServe(":8080", nil)
}
