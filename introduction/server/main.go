package main

import "net/http"

func main() {
	http.HandleFunc("/confirm", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Go webserver confirmed"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
