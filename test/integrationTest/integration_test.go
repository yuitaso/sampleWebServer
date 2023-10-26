package integrationtest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"testing"

	"gotest.tools/v3/assert"
)

var httpClient = new(http.Client)
var urlBase = "http://localhost:8080"

func TestPingRouter(t *testing.T) {
	testUrl := urlBase + "/internal/ping"
	req, _ := http.NewRequest("GET", testUrl, nil)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := io.ReadAll(resp.Body)
	var jsonRes map[string]interface{}
	_ = json.Unmarshal(byteArray, &jsonRes)

	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, jsonRes["message"].(string), "ok")
}

func TestCreateUser(t *testing.T) {
	requestEmail := "sample@hoge.com"
	requestPassword := "xxx"

	testUrl := urlBase + "/user/create"
	form := url.Values{}
	form.Add("email", requestEmail)
	form.Add("password", requestPassword)

	resp, err := httpClient.PostForm(testUrl, form)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := io.ReadAll(resp.Body)
	var jsonRes map[string]interface{}
	_ = json.Unmarshal(byteArray, &jsonRes)

	assert.Equal(t, resp.StatusCode, 200)
}
