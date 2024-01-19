package voicerss

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func ConvertToSpeech(language string, text string) []byte {

	voicerssKey := os.Getenv("VOICERSS_KEY")
	rapidapiKey := os.Getenv("RAPIDAPI_KEY")

	// call text-to-speech api
	url := "https://voicerss-text-to-speech.p.rapidapi.com/?key=" + voicerssKey

	payload := strings.NewReader("src=" + text + "&hl=" + language + "&c=wav&f=8khz_8bit_mono")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("x-rapidapi-host", "voicerss-text-to-speech.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", rapidapiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body
}
