package secretmanager

import (
	"encoding/json"
	"fmt"

	"twitter_go/awsgo"
	"twitter_go/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secretData models.Secret
	fmt.Println("> Ask secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	pass, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*pass.SecretString), &secretData)
	fmt.Println("> Secret read OK " + secretName)
	return secretData, nil
}
