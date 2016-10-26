package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getUser() (User, error) {
	endpoint := server + "/v2/users/me"
	var user User
	err := getRequest(endpoint, &user)
	return user, err
}

func getRequest(url string, result interface{}) error {
	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, result)
}

func postRequest() {

}
