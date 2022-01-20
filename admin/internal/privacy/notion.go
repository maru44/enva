package privacy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type (
	notionResponse struct {
		Results []notionResult `json:"results"`
		HasMore bool           `json:"has_more"`
	}

	notionResult struct {
		Properties notionProps `json:"properties"`
	}

	notionProps struct {
		Name struct {
			Title []struct {
				PlainText string `json:"plain_text"`
			} `json:"title"`
		} `json:"Name"`
		ConEn struct {
			RichText []struct {
				Text struct {
					Content string `json:"content"`
				} `json:"text"`
				PlainText string `json:"plain_text"`
			} `json:"rich_text"`
		} `json:"Content"`
	}
)

func getByAPI(token, notionDBID string) (*privacy, error) {
	var contents []string

	url := fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", notionDBID)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Notion-Version", "2021-08-16")
	req.Header.Set("Content-Type", "application/json")

	cli := http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("failed to request")
	}
	var data notionResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	for _, d := range data.Results {
		if d.Properties.ConEn.RichText == nil || len(d.Properties.ConEn.RichText) == 0 {
			continue
		}
		content := d.Properties.ConEn.RichText[0].Text.Content
		for _, r := range replacer {
			content = strings.ReplaceAll(content, fmt.Sprintf("[%s]", r.Signal), r.To)
		}
		contents = append(contents, content)
	}

	return &privacy{
		Contents: contents,
		Date:     time.Now().Format("Jan 2, 2006"),
	}, nil
}
