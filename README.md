# words-occurence
Takes a string/text and returns the top 10 occuring words in json format

How to use:
- clone the project on your local
- go build WordsOccurence
- go run WordsOccurence
- or simple use your ide run functionality
- server starts at 8085
- make a call to http://localhost:8085/frequency
  - example: curl -d "text=this is      is it police polo foli is      foli oh oh oh jio oh kolo polo plo soloo   soloo harvey harvey folo fou folo fiii oh&fetchAll=false" -X POST http://localhost:8085/frequency
  - or just use postman; your could use {"text": "this is      is it police polo foli is      foli oh oh oh jio oh kolo polo plo soloo   soloo harvey harvey folo fou folo fiii oh", "fetchAll": "true"} as raw input
  - text param: the actual text to process
  - fetchAll params: true= return all words with occurences not just the top 10; false= return only the top ten
