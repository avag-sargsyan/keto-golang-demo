package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	oryBaseURL = "https://flamboyant-keller-zdtbkurucf.projects.oryapis.com"
	apiKey     = "ory_pat_dLxOTPgdWVYWp1pSTKcjcMwRYRBA1hBy"
)

type CheckResponse struct {
	Allowed bool `json:"allowed"`
}

func checkPermission(namespace, object, relation, subject string) bool {
	data := fmt.Sprintf(`{
		"namespace": "%s",
		"object": "%s",
		"relation": "%s",
		"subject_id": "%s"
	}`, namespace, object, relation, subject)

	req, _ := http.NewRequest("POST", oryBaseURL+"/relation-tuples/check", bytes.NewBuffer([]byte(data)))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Request error: %v", err)
		return false
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Println("Keto API Response:", string(body))

	var checkResp CheckResponse
	if err := json.Unmarshal(body, &checkResp); err != nil {
		log.Printf("Decode error: %v", err)
		return false
	}

	return checkResp.Allowed
}

func main() {
	if checkPermission("Order", "111", "owner", "alice") {
		fmt.Println("✅ Alice owns order:111")
	} else {
		fmt.Println("❌ Alice does not own order:111")
	}
}
