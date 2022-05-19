package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"speech/helpers"
	"speech/model"
)

func main() {
	requestURL := "https://texttospeech.googleapis.com/v1beta1/text:synthesize"
	apiKey := "Google_API_Key"

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter a text: ")
	scanner.Scan()

	d := model.RequestBody{
		AudioConfig: model.AudioConfig{
			AudioEncoding: "LINEAR16",
			Pitch:         1,
			SpeakingRate:  1,
		},
		Input: map[string]string{
			"text": scanner.Text(),
		},
		Voice: map[string]string{
			"languageCode": "en-US",
			"name":         "en-US-Wavenet-E",
		},
	}

	data, err := json.Marshal(d)
	helpers.CheckError(err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(data))
	helpers.CheckError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", apiKey)

	fmt.Println("Processing...")

	resp, err := client.Do(req)
	helpers.CheckError(err)

	body, err := io.ReadAll(resp.Body)
	helpers.CheckError(err)
	helpers.CheckError(resp.Body.Close())

	result := model.ResponseBody{}
	err = json.Unmarshal(body, &result)
	helpers.CheckError(err)

	err = os.WriteFile("mp3/test.mp3", result.AudioContent, 0644)
	helpers.CheckError(err)
	fmt.Println("test.mp3 file created")
}
