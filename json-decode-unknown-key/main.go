package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
// some ElasticSearch reading
settings: {
	index: {
		creation_date: "1458940066227",
		number_of_shards: "10",
		number_of_replicas: "2",
		uuid: "srPAADNKTjKjbDSsbOyUBw",
		version: {
			created: "2020099"
		}
	}
*/

type Index struct {
	Settings Settings
}

type Settings struct {
	IndexDef IndexDef `json:"index"`
}

type IndexDef struct {
	CreateDate       string `json:"creation_date"`
	NumberOfShards   string `json:"number_of_shards"`
	NumberOfReplicas string `json:"number_of_replicas"`
	Uuid             string `json:"uuid"`
}

func main() {
	addr := flag.String("address", "localhost:9200", "host:port pair to your ElasticSearch cluster.")
	flag.Parse()
	response, err := http.Get(fmt.Sprintf("http://%v/_settings", *addr))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var v map[string]Index
	json.Unmarshal(body, &v)
	fmt.Println(v)

	for key, settings := range v {
		fmt.Printf("%v => %+v\n", key, settings)
	}
}
