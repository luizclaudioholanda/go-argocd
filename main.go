package main

import "net/http"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world from GO!"))
	})
	http.ListenAndServe(":9000", nil)

}
