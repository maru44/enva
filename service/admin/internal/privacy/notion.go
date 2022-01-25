package privacy

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/maru44/enva/service/api/pkg/tools"
)

type (
	notionResponse struct {
		Results    []notionResult `json:"results"`
		HasMore    bool           `json:"has_more"`
		NextCursor *string        `json:"next_cursor"`
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
		Num struct {
			Number int `json:"number"`
		} `json:"Num"`
	}

	notionSort struct {
		Property  string `json:"property"`
		Direction string `json:"direction"`
	}

	notionRequestBody struct {
		StartCursor *string      `json:"start_cursor,omitempty"`
		PageSize    int32        `json:"page_size,omitempty"`
		Sorts       []notionSort `json:"sorts,omitempty"`
	}
)

func getByAPI(token, notionDBID string) (*privacy, error) {
	var (
		contents    []string
		startCursor string
		url         = fmt.Sprintf("https://api.notion.com/v1/databases/%s/query", notionDBID)
	)

	for {
		input := notionRequestBody{
			StartCursor: tools.StringPtr(startCursor),
			PageSize:    100,
			Sorts: []notionSort{
				{
					Property:  "Num",
					Direction: "ascending",
				},
			},
		}
		inputJ, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(inputJ))
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

		if !data.HasMore || data.NextCursor == nil {
			break
		}
		startCursor = *data.NextCursor
	}

	return &privacy{
		Contents: contents,
		Date:     time.Now().Format("Jan 2, 2006"),
	}, nil
}
