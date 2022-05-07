package tiktok_tts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Voice string

const (
	VoiceEnglishFem1 = "en_us_001"
)

const URL = "https://api16-normal-useast5.us.tiktokv.com/media/api/text/speech/invoke/?text_speaker=%s&req_text=%s&speaker_map_type=0"

type TTSResponse struct {
	Data       Data   `json:"data"`
	Message    string `json:"message"`
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Data struct {
	SKey     string `json:"s_key"`
	VStr     string `json:"v_str"`
	Duration string `json:"duration"`
}

func TTS(voice Voice, text string) (TTSResponse, error) {
	text = strings.ReplaceAll(text, "+", "plus")
	text = strings.ReplaceAll(text, " ", "+")
	text = strings.ReplaceAll(text, "&", "and")

	requestUrl := fmt.Sprintf(URL, voice, text)
	resp, err := http.Post(requestUrl, "text/plain", nil)
	if err != nil {
		return TTSResponse{}, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return TTSResponse{}, err
	}

	var r TTSResponse
	err = json.Unmarshal(respBytes, &r)
	return r, err
}
