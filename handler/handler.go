package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	url, apiKey := os.Getenv("URL"), os.Getenv("API_KEY")
	// Get query  ex: /search?s=Batman&page=2
	s := req.URL.Query().Get("s")
	page := req.URL.Query().Get("page")
	// Get response from URL
	res, err := http.Get(fmt.Sprintf("%v?apikey=%v&s=%v&page=%v", url, apiKey, s, page))
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write(resBody)
}

func DetailHandler(w http.ResponseWriter, req *http.Request) {
	detailId := mux.Vars(req)["id"]

	url, apiKey := os.Getenv("URL"), os.Getenv("API_KEY")

	// Get response from URL
	res, err := http.Get(fmt.Sprintf("%v?apikey=%v&i=%v", url, apiKey, detailId))
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		w.Write([]byte("error"))
		return
	}
	w.Write(resBody)
}
