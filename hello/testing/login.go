package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://202.6.215.230:9003/rest/login"

	payload := strings.NewReader("{\n  \"user_id\" : \"U000007\",\n  \"password\" : \"Berta1!\",\n  \"app_id\" : \"BERTA\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "0168046e-c932-4bd9-bf1c-18eaddd42830")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// fmt.Println(res)
	fmt.Println(string(body))

}
