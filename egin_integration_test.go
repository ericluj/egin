package egin

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testRequest(t *testing.T, params ...string) {
	if len(params) == 0 {
		t.Fatal("url cannot be empty")
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(params[0])
	assert.NoError(t, err)
	defer resp.Body.Close()

	body, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)

	var responseStatus = "200 OK"
	if len(params) > 1 && params[1] != "" {
		responseStatus = params[1]
	}

	var responseBody = "it worked"
	if len(params) > 2 && params[2] != "" {
		responseBody = params[2]
	}

	assert.Equal(t, responseStatus, resp.Status, "should get a "+responseStatus)
	if responseStatus == "200 OK" {
		assert.Equal(t, responseBody, string(body), "resp body should match")
	}
}

func TestRunEmpty(t *testing.T) {
	router := New()
	go func() {
		router.GET("/example", func(c *Context) {})
		assert.NoError(t, router.Run())
	}()

	time.Sleep(5 * time.Millisecond)

	assert.Error(t, router.Run(":8080"))
	testRequest(t, "http://localhost:8080/example")
}
