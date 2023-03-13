package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"wikipediaTest/src/entities"
)

func WikipediaApi(request string) (answer []string) {
	s := make([]string, 3)
	if response, err := http.Get(request); err != nil {
		s[0] = "Wikipedia is not respond"
	} else {
		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		sr := &entities.SearchResults{}
		if err = json.Unmarshal(contents, sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		if !sr.Ready {
			s[0] = "Something going wrong, try to change your question"
		}

		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}
	return s
}
