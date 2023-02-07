package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}

func Create(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Printf("invalid request")
	}

	fmt.Printf("Name of user %v", user.Name)
}
