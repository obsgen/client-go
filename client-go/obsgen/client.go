package obsgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type ObsGenClient struct {
	apiKey   string
	baseID   string
	table    string
	client   *http.Client
	endpoint string
}

func NewClient(apiKey string) (*ObsGenClient, error) {
	parts := strings.Split(apiKey, "-")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid apiKey format: %s", apiKey)
	}

	apiKeyPart := parts[0]
	baseID := parts[1]
	table := parts[2]

	endpoint := fmt.Sprintf("https://api.airtable.com/v0/%s/%s", baseID, table)
	client := &http.Client{}

	return &ObsGenClient{
		apiKey:   apiKeyPart,
		baseID:   baseID,
		table:    table,
		client:   client,
		endpoint: endpoint,
	}, nil
}

func (c *ObsGenClient) LogEvent(data map[string]interface{}) error {
	requestBody := map[string]interface{}{
		"records": []map[string]interface{}{
			{"fields": data},
		},
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(requestJSON))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	request.Header.Set("Content-Type", "application/json")

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("unexpected response status code: %d", response.StatusCode)
	}

	return nil
}