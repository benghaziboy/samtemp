package samtemp

import (
	"encoding/csv"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/exp/ses"
	"log"
	"os"
)

var (
	Ses *ses.SES

	AccessKey        = os.Getenv("AWS_ACCESS_KEY")
	SecretKey        = os.Getenv("AWS_SECRET_KEY")
	AwsCredsFilepath = os.Getenv("AWS_CREDS_FILEPATH")
)

func authenticate() (aws.Auth, error) {
	var auth aws.Auth

	if AccessKey != "" && SecretKey != "" {
		auth, err := aws.EnvAuth()
		if err != nil {
			return auth, err
		}

		return auth, nil
	}

	if AwsCredsFilepath == "" {
		AwsCredsFilepath = "credentials.csv"
	}

	file, err := os.Open(AwsCredsFilepath)
	if err != nil {
		return auth, err
	}

	reader := csv.NewReader(file)
	csvData, err := reader.ReadAll()
	if err != nil {
		return auth, err
	}

	auth.AccessKey = csvData[1][1]
	auth.SecretKey = csvData[1][2]

	return auth, nil
}

func init() {
	auth, err := authenticate()
	if err != nil {
		log.Println(err.Error())
	}

	Ses = ses.New(auth, aws.USEast)
	if err != nil {
		log.Println(err)
	}
}
