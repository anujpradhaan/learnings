package resttemplate

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func MakeGetCall(url string, response interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, response)
}
