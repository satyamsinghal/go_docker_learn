package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age string `json:"age" validate:"required"`
	Profession string `json:"profession" validate:"required"`
}

func Create(w http.ResponseWriter, req *http.Request) {
	bytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("invalid request err: %v", err)
	}
	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Printf("invalid request err: %v", err)
	}

	//validator.Validate.

	fmt.Printf("Name of user %v", user.Name)
	w.WriteHeader(http.StatusNoContent)
}
