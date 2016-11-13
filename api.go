package monobullet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getUser() (User, error) {
	var user User
	err := getRequest(meEndpoint, &user)
	return user, err
}

func getRequest(endpoint string, result interface{}) error {
	client := &http.Client{}
	u := url.URL{Scheme: "https", Host: apiServer, Path: endpoint}
	resp, err := client.Get(u.String())
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
