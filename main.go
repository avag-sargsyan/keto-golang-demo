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

func permissionHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	namespace := query.Get("namespace")
	object := query.Get("object")
	relation := query.Get("relation")
	subject := query.Get("subject")

	if namespace == "" || object == "" || relation == "" || subject == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	allowed := checkPermission(namespace, object, relation, subject)
	response := map[string]bool{"allowed": allowed}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/check-permission", permissionHandler)

	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
