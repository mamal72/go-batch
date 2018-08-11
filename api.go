package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const batchTransactionalSendURLPattern = "https://api.batch.com/1.1/%s/transactional/send"

func getBatchTransactionalSendURL(batchAPIKey string) string {
	return fmt.Sprintf(batchTransactionalSendURLPattern, batchAPIKey)
}

// NewClient returns a new instance of batch api client
func NewClient(restAPIKey, batchAPIKey string) APIClient {
	return APIClient{
		RESTAPIKey:  restAPIKey,
		BatchAPIKey: batchAPIKey,
	}
}

// TransactionalSend sends a transactional push notification and
// returns token in case of success
// See https://batch.com/doc/api/transactional.html
func (b *APIClient) TransactionalSend(payload TransactionalPushPayload) (TransactionalPushResponse, error) {
	pushResponse := TransactionalPushResponse{}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return pushResponse, err
	}
	payloadBody := bytes.NewReader(payloadBytes)

	batchAPIURL := getBatchTransactionalSendURL(b.BatchAPIKey)
	req, err := http.NewRequest("POST", batchAPIURL, payloadBody)
	if err != nil {
		return pushResponse, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Authorization", b.RESTAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return pushResponse, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pushResponse, err
	}

	if resp.StatusCode != http.StatusCreated {
		return pushResponse, fmt.Errorf("invalid batch response: %d - %s", resp.StatusCode, string(bodyBytes))
	}

	err = json.Unmarshal(bodyBytes, &pushResponse)
	return pushResponse, err
}
