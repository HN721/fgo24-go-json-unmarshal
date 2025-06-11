package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Datas struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Company Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
}

func Service() {
	url := "https://jsonplaceholder.typicode.com/users"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Datas []Datas
	err = json.Unmarshal(body, &Datas)
	if err != nil {
		log.Fatal(err)
	}
	for _, Datas := range Datas {
		fmt.Printf("Name: %s\nEmail: %s\nCompany: %s\n\n", strings.ToLower(Datas.Name), strings.ToLower(Datas.Email), strings.ToLower(Datas.Company.Name))
	}
}
