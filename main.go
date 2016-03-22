// License: MIT
package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataMaster struct {
	ListenAddr string
	ListenPort uint
	db         *sql.DB
}

func (datamaster *DataMaster) Run() {
	var (
		err            error
		dataSourceName string
	)

	dataSourceName = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8",
		"root", "", "127.0.0.1:3306", "sqltap_development")

	datamaster.db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to DB. %v\n", err)
	}

	err = datamaster.db.Ping()
	if err != nil {
		log.Fatalf("cannot ping to db. %v\n", err)
	}

	http.HandleFunc("/", datamaster.helloWorld)
	http.HandleFunc("/fun/with/channels", datamaster.funWithChannels)

	addr := fmt.Sprintf("%v:%v", datamaster.ListenAddr, datamaster.ListenPort)
	log.Printf("Listening on %v\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Failed to listen. %v\n", err)
	}
}

type CleverResponse struct {
	Error      error
	Url        string
	Body       string
	StatusCode int
}

func (datamaster *DataMaster) funWithChannels(w http.ResponseWriter, r *http.Request) {
	var (
		urls = []string{
			"http://en.dawanda.com/",
			"https://www.google.com/",
			"https://www.bsi.bund.de/DE/Home/home_node.html",
		}
		responseChannel chan CleverResponse = make(chan CleverResponse)
	)

	// initiate distributed GET requests
	for _, url := range urls {
		go datamaster.HttpGet(url, responseChannel)
	}

	// consumer
	var timeout = time.After(time.Second * 5)
	var i int
	for {
		select {
		case response := <-responseChannel:
			i++
			fmt.Fprintf(w, "response %v received. status:%v contentLength:%v\n", i, response.StatusCode, len(response.Body))
			if i == len(urls) {
				fmt.Fprintf(w, "This was the last one.\n")
				return
			}
		case <-timeout:
			w.WriteHeader(http.StatusGatewayTimeout)
			fmt.Fprintf(w, "Time's up!\n")
			return
		}
	}
}

func (datamaster *DataMaster) HttpGet(url string, responseChannel chan CleverResponse) {
	response, err := http.Get(url)
	if err != nil {
		responseChannel <- CleverResponse{Error: err}
	} else {
		defer response.Body.Close()
		bytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			responseChannel <- CleverResponse{Error: err}
		} else {
			responseChannel <- CleverResponse{
				Url:        url,
				StatusCode: response.StatusCode,
				Body:       string(bytes),
			}
		}
	}
}

func (datamaster *DataMaster) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/cheers")

	target := r.URL.Query().Get("target")

	var id int
	var email string
	var err error
	err = datamaster.db.QueryRow("SELECT id, email FROM users WHERE username = ?", target).Scan(&id, &email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB SELECT failed. %v\n", err)
		return
	}

	fmt.Fprintf(w, "id %v, email %v\n", id, email)
}

func main() {
	var datamaster = DataMaster{
		ListenAddr: "0.0.0.0",
		ListenPort: 8081,
	}

	datamaster.Run()
}
