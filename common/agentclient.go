package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ExplainResponse struct {
	Topic      string   `json:"topic"`
	Summary    string   `json:"summary"`
	Suggestion string   `json:"suggestion"`
	Sources    []string `json:"sources"`
	ToolsUsed  []string `json:"tools_used"`
}

type AgentResponse struct {
	Message string `json:"message"`
}

type EmbeddingResponse struct {
	Embedding []any  `json:"embedding"`
	Error     string `json:"error"`
}

func GetAgentBaseUrl() string {
	defaultUrl := "http://localhost:8000/"
	url, exists := os.LookupEnv("CHATCTL_AGENT_BASE_URL")
	if !exists {
		return defaultUrl
	}
	return url
}

func GenerateEmbedding(model string, text string) (*EmbeddingResponse, error) {
	url := GetAgentBaseUrl()

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "embedding"
	payload := strings.NewReader(`{"text":"` + text + `","model":"` + model + `"}`)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling embedding endpoint: %s", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response EmbeddingResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing init response: %s", err)
	}
	return &response, nil
}

func InitAgent(model string) (*AgentResponse, error) {
	url := GetAgentBaseUrl()

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "init"
	payload := strings.NewReader(`{"model":"` + model + `"}`)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("CHATCTL_OPEN_API_KEY"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling init endpoint: %s", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response AgentResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing init response: %s", err)
	}
	return &response, nil

}

func CallAgent(query string) (*ExplainResponse, error) {
	url := GetAgentBaseUrl()
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += "explain"
	payload := strings.NewReader(`{"text":"` + query + `"}`)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling explain agent: %s", err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var response ExplainResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error parsing explain response: %s", err)
	}
	return &response, nil
}
