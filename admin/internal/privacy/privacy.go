package privacy

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (
	privacy struct {
		Content string `json:"content"`
	}

	replace struct {
		Signal string
		To     string
	}
)

const privacyJson = "./service/front/src/components/static/rule/privacy.json"

var replacer = []replace{
	{
		Signal: "date",
		To:     time.Now().Format("Jan 2, 2006"),
	},
}

func GenPrivacyJson() error {
	notionDBID := os.Getenv("N_PRIVACY_TABLE")
	notionToken := os.Getenv("N_READ_TOKEN")

	if notionDBID == "" || notionToken == "" {
		fmt.Println("tokens not set")
		return nil
	}

	ps, err := getByAPI(notionToken, notionDBID)
	if err != nil {
		return err
	}

	j, err := json.Marshal(ps)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(privacyJson, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(j); err != nil {
		return err
	}
	return nil
}
