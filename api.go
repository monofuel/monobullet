package monobullet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func getUser() (*User, error) {
	user := new(User)
	err := getRequest(meEndpoint, user, "")
	return user, err
}

type GetPushParams struct {
	ModifiedAfter int32
	Active        bool
	Cursor        string
	Limit         int
}
type getResp struct {
	Pushes        []*Push       `json:"pushes"`
	Profiles      []interface{} `json:"profiles"`
	Subscriptions []interface{} `json:"subscriptions"`
	Blocks        []interface{} `json:"blocks"`
	Chats         []interface{} `json:"chats"`
	Contacts      []interface{} `json:"contacts"`
	Devices       []interface{} `json:"devices"`
	Grants        []interface{} `json:"grants"`
	Accounts      []interface{} `json:"accounts"`
	Channels      []interface{} `json:"channels"`
	Clients       []interface{} `json:"clients"`
	Texts         []interface{} `json:"texts"`
}

func getPushes(params GetPushParams) ([]*Push, error) {

	var pushes getResp
	var query []string
	if params.ModifiedAfter != 0 {
		query = append(query, fmt.Sprintf("modified_after=%v", params.ModifiedAfter))
	}
	if params.Active {
		query = append(query, "active=true")
	}
	if params.Cursor != "" {
		query = append(query, fmt.Sprintf("cursor=%v", params.Cursor))
	}
	if params.Limit != 0 {
		query = append(query, fmt.Sprintf("limit=%v", params.Limit))
	}

	err := getRequest(pushEndpoint, &pushes, strings.Join(query, "&"))
	return pushes.Pushes, err
}

func sendPush(payload *Push) (*Push, error) {
	resp := new(Push)
	err := postRequest(pushEndpoint, payload, resp)
	return resp, err
}

func getRequest(endpoint string, result interface{}, urlParams string) error {
	client := &http.Client{}
	u := url.URL{Scheme: "https", Host: apiServer, Path: endpoint, RawQuery: urlParams}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Access-Token", config.APIKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		if config.Debug {
			fmt.Printf("bad response for GET %v | %v\n", u.String(), string(body))
		}
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	if config.Debug {
		fmt.Printf("response for GET %v | %v\n", u.String(), string(body))
		findUnusedFields(body, result)
	}

	return json.Unmarshal(body, result)
}

func postRequest(endpoint string, payload interface{}, result interface{}) error {
	client := &http.Client{}
	u := url.URL{Scheme: "https", Host: apiServer, Path: endpoint}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	req.Header.Add("Access-Token", config.APIKey)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		if config.Debug {
			fmt.Printf("bad response for POST %v | %v\n", u.String(), string(body))
		}
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	if config.Debug {
		fmt.Printf("response for POST %v | %v\n", u.String(), string(body))
		findUnusedFields(body, result)
	}

	return json.Unmarshal(body, result)
}

// note: does not handle extended fields
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
