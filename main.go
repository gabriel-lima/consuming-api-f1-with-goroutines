package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Circuit struct {
	CircuitName string `json:"circuitName"`
}

func worker(id int, jobs <-chan func() []byte, results chan<- []byte) {
	for job := range jobs {
		results <- job()
	}
}

func main() {

	const numberOfCircuits = 20
	const numberOfJobs = 1
	const numberOfWorks = 1

	jobs := make(chan func() []byte, numberOfJobs)
	results := make(chan []byte, numberOfCircuits)

	function := func(id int) func() []byte {
		return func() []byte {

			url := "http://ergast.com/api/f1/current/" + strconv.Itoa(id) + ".json"
			response, _ := http.Get(url)
			defer response.Body.Close()

			body, _ := ioutil.ReadAll(response.Body)
			return body
		}
	}

	for w := 1; w <= numberOfWorks; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numberOfCircuits; j++ {
		jobs <- function(j)
	}
	close(jobs)

	for a := 1; a <= numberOfCircuits; a++ {
		var circuit Circuit
		var rawResults map[string]map[string]map[string][]map[string]interface{}
		json.Unmarshal(<-results, &rawResults)
		rawCircuit, _ := json.Marshal(rawResults["MRData"]["RaceTable"]["Races"][0]["Circuit"])
		json.Unmarshal(rawCircuit, &circuit)
		fmt.Println(circuit)
	}
}
