package monobullet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

func getUser() (User, error) {
	var user User
	err := getRequest(meEndpoint, &user)
	return user, err
}

func getRequest(endpoint string, result interface{}) error {
	client := &http.Client{}
	u := url.URL{Scheme: "https", Host: apiServer, Path: endpoint}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Access-Token", config.APIKey)
	resp, err := client.Do(req)
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
	if config.Debug {
		fmt.Printf("response for GET %v | %v\n", u.String(), string(body))
		findUnusedFields(body, result)
	}

	return json.Unmarshal(body, result)
}

func postRequest() {

}

func findUnusedFields(buf []byte, result interface{}) {
	var all map[string]interface{}
	err := json.Unmarshal(buf, &all)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(buf, result)
	if err != nil {
		log.Fatal(err)
	}
	resType := reflect.TypeOf(result).Elem()
	for key := range all {
		found := false
		for i := 0; i < resType.NumField(); i++ {
			field := resType.Field(i)
			tag := field.Tag
			if tag.Get("json") == key {
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("field missing: %v\n", key)
		}
	}
}
