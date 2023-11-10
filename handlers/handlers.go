package handlers

import (
	"context"
	"fmt"

	"twitter_go/models"

	"twitter_go/jwt"
	"twitter_go/routers"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status = 400

	isOK, statusCode, msg, _ := validateAuthorization(ctx, request)
	if !isOK {
		r.Status = statusCode
		r.Message = msg
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "signup":
			return routers.Signup(ctx)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Method Invalid"
	return r
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "signup" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token required", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !todoOK {
		if err != nil {
			fmt.Println("Error in the token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error in the token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim
}
