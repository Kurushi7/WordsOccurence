package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Text      string `json:"text"`
	Occurence int    `json:"occurence"`
}

type PairList []Pair

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// parse the raw query
		if err := r.ParseForm(); err != nil {
			sendJsonEncodedText(w, http.StatusBadRequest, "ParseForm() err: "+err.Error())
		}

		// the actual text for processing
		text := r.Form.Get("text")

		// do we filter the top 10 results or not;
		// true:  fetch all
		// false: return the top ten
		fetchAll := r.Form.Get("fetchAll")
		boolValue, _ := strconv.ParseBool(fetchAll)

		result := processString(text, boolValue)

		sendJsonEncodedText(w, http.StatusOK, result)
	case "GET":
		sendJsonEncodedText(w, http.StatusBadRequest, "Only post action is supported")
	default:
		sendJsonEncodedText(w, http.StatusBadRequest, "Only post action is supported")
	}
}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

/*
	Function takes a string and returns a map of words and their respective occurences in the string
*/
func wordCount(str string) map[string]int {
	words := strings.Fields(str)
	wordMap := make(map[string]int)

	for _, word := range words {
		_, ok := wordMap[word]
		if ok {
			wordMap[word] += 1
		} else {
			wordMap[word] = 1
		}
	}
	return wordMap
}

/*
	params:
		str: the string being passed for manipulation
		fetchAll: says if we return the top ten only or the whole set of result
	This function takes a string and returns the word frequency of each word in the string or the top ten depending on fetchAll
*/
func processString(str string, fetchAll bool) PairList {

	strMap := wordCount(str)

	// build an array of struct out of the map list
	keys := make(PairList, len(strMap))
	i := 0
	for index, value := range strMap {
		keys[i] = Pair{index, value}
		i++
	}

	// sort the struct in a descending fashion
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Occurence > keys[j].Occurence
	})

	if !fetchAll {
		if len(keys) > 10 {
			return keys[:10]
		}
	}
	return keys

}

func sendJsonEncodedText(w http.ResponseWriter, statusCode int, results interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		sendJsonEncodedText(w, http.StatusBadRequest, "Error while writing to stream: "+err.Error())
	}
}
