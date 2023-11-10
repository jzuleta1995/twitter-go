package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"twitter_go/bd"
	"twitter_go/models"
)

func Signup(ctx context.Context) models.RespApi {
	var t models.User
	var r models.RespApi
	r.Status = 400

	fmt.Println("Start signup")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Must specify the Email"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Must specify a password with at least 6 characters"
		fmt.Println(r.Message)
		return r
	}

	_, found, _ := bd.UserExists(t.Email)
	if found {
		r.Message = "An User with the Email " + t.Email + " already exists "
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertUser(t)
	if err != nil {
		r.Message = "An error has occurred while storing the user " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "The user could not be stored"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Register OK"
	fmt.Println(r.Message)
	return r
}
