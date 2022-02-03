package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUrlPath(t *testing.T) {
	r := strings.NewReader(`{"text": "this is      is it police polo foli is      foli oh oh oh jio oh kolo polo plo soloo   soloo harvey harvey folo fou folo fiii oh", "fetchAll": true}`)
	req, err := http.NewRequest("POST", "http://localhosts:8085/testo", r)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleStringProcessing(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(resBody))
	assert.Equal(t, string(resBody), "404 not found.\n")
	assert.Equal(t, rr.Code, http.StatusNotFound)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestMethod(t *testing.T) {
	r := strings.NewReader(`{"text": "this is      is it police polo foli is      foli oh oh oh jio oh kolo polo plo soloo   soloo harvey harvey folo fou folo fiii oh", "fetchAll": true}`)
	req, err := http.NewRequest("GET", "http://localhosts:8085/frequency", r)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handleStringProcessing(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, string(resBody), "\"Only post action is supported\"\n")
	assert.Equal(t, rr.Code, http.StatusBadRequest)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestWordCount(t *testing.T) {

	testString := "this is it it it is isss for"

	stringMap := wordCount(testString)

	expectedStringMap := map[string]int{"for": 1, "is": 2, "isss": 1, "it": 3, "this": 1}

	assert.Equal(t, expectedStringMap, stringMap)

}

//func TestProcessString(t *testing.T) {
//
//	testString := "this is it it it is isss for folio sola"
//
//	stringMap := processString(testString, false)
//
//	var keys PairList
//	expectedPairList := append(keys, Pair{"it", 3}, Pair{"is", 2}, Pair{Text: "for", Occurence: 1}, Pair{"folio", 1}, Pair{"sola", 1}, Pair{"this", 1}, Pair{"isss", 1})
//
//	// sort the struct in a descending fashion
//	sort.Slice(expectedPairList, func(i, j int) bool {
//		return expectedPairList[i].Occurence > expectedPairList[j].Occurence
//	})
//
//	assert.Equal(t, expectedPairList, stringMap)
//
//}
