package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type AskResponse struct {
	Topic     string   `json:"topic"`
	Summary   string   `json:"summary"`
	Sources   []string `json:"sources"`
	ToolsUsed []string `json:"tools_used"`
}

func CallAgent(query string) (*AskResponse, error) {
	url := os.Getenv("CHATCTL_AGENT_BASE_URL")
	// add trailing slash if not present
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "ask"
	payload := strings.NewReader(`{"query":"` + query + `"}`)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auhthorization", "Bearer "+os.Getenv("CHATCTL_OPEN_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling Agent: %s", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response AskResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing Agent response: %s", err)
	}
	return &response, nil
}
