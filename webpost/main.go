package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rodrigobotti/go-lessons/webpost/model"
)

var (
	userAgent = "Fala galeris. Botti aqui. Testando consumo de apis rest com go-lang. Relax, foi chamada besta de casa sem carga."
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 1,
	}

	user := model.LoginRequest{
		Phone: "11940036805",
		Token: "111111",
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("[main] error creating json from user")
	}

	request, err := http.NewRequest(http.MethodPost, "https://mobileapp-v2.nextel.com.br/auth/v.3.0/login", bytes.NewBuffer(userJSON))
	if err != nil {
		fmt.Println("[main] error creating post request", err.Error())
		return
	}
	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("[main] error posting to api", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("[main] error reading api response", err.Error())
			return
		}
		login := model.LoginResponse{}
		err = json.Unmarshal(body, &login)
		if err != nil {
			fmt.Println("[main] error when desserializing response", err.Error())
		}
		fmt.Println("-- login -- \n", login)
	} else {
		fmt.Println("[main] api returned status = ", res.StatusCode)
		return
	}
}
