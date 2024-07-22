package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int `json:"age"`
}

func unMarshall() {
	jsonString := `{
		"name": "Daniel Ronaldo Pangestu",
		"email": "ronaldo.pangestu1@gmail.com",
		"age": 23
	}`

	var jsonData = []byte(jsonString)

	var user User

	var err = json.Unmarshal(jsonData, &user)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name: ", user.Name)
	fmt.Println("Email: ", user.Email)
	fmt.Println("Age: ", user.Age)
}

func marshal() {
	user := User{Name: "Daniel", Email: "ronaldo.pangestu1@gmail.com", Age: 30}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(jsonData))
}

func main() {
	unMarshall()
	marshal()
}