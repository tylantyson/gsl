package gsl

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Storage
type Storage struct {
	accessKey string
	endpoint  string
	zoneName  string
}

// Delete
func (storage *Storage) Delete(path string) (bool, error) {
	// Create HTTP Request
	request, err := http.NewRequest("DELETE", fmt.Sprintf("https://%s/%s/%s", storage.endpoint, storage.zoneName, path), nil)
	// Handle Error
	if err != nil {
		return false, err
	}
	// Add Headers
	request.Header.Add("AccessKey", storage.accessKey)
	// Do Request
	response, err := http.DefaultClient.Do(request)
	// Handle Error
	if err != nil {
		return false, err
	}
	// Close Response
	defer response.Body.Close()
	// Check Respnse
	if response.StatusCode/100 != 2 {
		// Ready Body
		body, err := io.ReadAll(response.Body)
		// Handle Error
		if err != nil {
			return false, err
		}
		// Response Data
		responseData := struct {
			Message string `json:"Message"`
		}{}
		// Unmarshal Body
		err = json.Unmarshal(body, &responseData)
		// Handle Error
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("%d : %s", response.StatusCode, responseData.Message)
	}
	// Success
	return true, err
}

// Download
func (storage *Storage) Download(path string) ([]byte, error) {
	// Create HTTP Request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s/%s/%s", storage.endpoint, storage.zoneName, path), nil)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Add Headers
	request.Header.Add("accept", "*/*")
	request.Header.Add("AccessKey", storage.accessKey)
	// Do Request
	response, err := http.DefaultClient.Do(request)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Close Response
	defer response.Body.Close()
	// Read Body
	body, err := io.ReadAll(response.Body)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Check Response
	if response.StatusCode/100 != 2 {
		// Response Data
		responseData := struct {
			Message string `json:"Message"`
		}{}
		// Unmarshal Body
		err := json.Unmarshal(body, &responseData)
		// Handle Error
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%d : %s", response.StatusCode, responseData.Message)
	}
	// Success
	return body, err
}

// List
func (storage *Storage) List(path string) ([]Object, error) {
	// Create HTTP Request
	request, err := http.NewRequest("GET", fmt.Sprintf("https://%s/%s/%s", storage.endpoint, storage.zoneName, path), nil)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Add Headers
	request.Header.Add("accept", "application/json")
	request.Header.Add("AccessKey", storage.accessKey)
	// Do Request
	response, err := http.DefaultClient.Do(request)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Close Response
	defer response.Body.Close()
	// Ready Body
	body, err := io.ReadAll(response.Body)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Check Response
	if response.StatusCode/100 != 2 {
		// Response Data
		responseData := struct {
			Message string `json:"Message"`
		}{}
		// Unmarshal Body
		err = json.Unmarshal(body, &responseData)
		// Handle Error
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%d : %s", response.StatusCode, responseData.Message)
	}
	// Response Data
	var responseData []Object
	// Unmarshal Body
	err = json.Unmarshal(body, &responseData)
	// Handle Error
	if err != nil {
		return nil, err
	}
	// Success
	return responseData, err
}

// Upload
func (storage *Storage) Upload(path string, data []byte) (bool, error) {
	// Create HTTP Request
	request, err := http.NewRequest("PUT", fmt.Sprintf("https:/%s/%s/%s", storage.endpoint, storage.zoneName, path), bytes.NewReader(data))
	// Handle Error
	if err != nil {
		return false, err
	}
	// Hasher
	hasher := sha256.New()
	hasher.Write(data)
	// Add Headers
	request.Header.Set("AccessKey", storage.accessKey)
	request.Header.Set("Checksum", strings.ToUpper(hex.EncodeToString(hasher.Sum(nil))))
	request.Header.Set("Content-Type", "application/octet-stream")
	// Do Request
	response, err := http.DefaultClient.Do(request)
	// Handle Error
	if err != nil {
		return false, err
	}
	// Close Response
	defer response.Body.Close()
	// Check Response
	if response.StatusCode/100 != 2 {
		// Read Body
		body, err := io.ReadAll(response.Body)
		// Handle Error
		if err != nil {
			return false, err
		}
		// Response Data
		responseData := struct {
			Message string `json:"Message"`
		}{}
		// Unmarshal Body
		err = json.Unmarshal(body, &responseData)
		// Handle Error
		if err != nil {
			return false, err
		}
		return false, fmt.Errorf("%d : %s", response.StatusCode, responseData.Message)
	}
	// Success
	return true, err
}
