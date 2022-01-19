package privacy

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type (
	Privacy struct {
		Content string `json:"content"`
	}

	Replace struct {
		Signal string
		To     string
	}
)

var Replacer = []Replace{
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

	fmt.Println(string(j))
	// @TODO write to json file

	return nil
}
