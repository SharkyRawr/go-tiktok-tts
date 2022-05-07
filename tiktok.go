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
	VoiceEnglishFem2 = "en_us_002"
	VoiceEnglishMal1 = "en_us_006"
	VoiceEnglishMal2 = "en_us_007"
	VoiceEnglishMal3 = "en_us_009"
	VoiceEnglishMal4 = "en_us_010"

	VoiceEnglishFemAU = "en_au_001"
	VoiceEnglishMalAU = "en_au_002"

	VoiceEnglishMalGB  = "en_uk_001"
	VoiceEnglishMalGB2 = "en_uk_003"

	VoiceFrenchMal1 = "fr_001"
	VoiceFrenchMal2 = "fr_002"

	VoiceGermanFem = "de_001"
	VoiceGermanMal = "de_002"

	VoiceEspMal = "es_002"

	VoiceSpaMal = "es_mx_002"

	VoiceBraFem1 = "br_001"
	VoiceBraFem2 = "br_003"
	VoiceBraFem3 = "br_004"
	VoiceBraMal  = "br_005"

	VoiceIdoFem = "id_001"

	VoiceJpnFem1 = "jp_001"
	VoiceJpnFem2 = "jp_003"
	VoiceJpnFem3 = "jp_005"
	VoiceJpnMal  = "jp_006"

	VoiceKorMal1 = "kr_002"
	VoiceKorFem  = "kr_003"
	VoiceKorMal2 = "kr_004"

	// Didney
	VoiceGhostface    = "en_us_ghostface"
	VoiceChewbacca    = "en_us_chewbacca"
	VoiceC3PO         = "en_us_c3po"
	VoiceStitch       = "en_us_stitch"
	VoiceStormtrooper = "en_us_stormtrooper"
	VoiceRocket       = "en_us_rocket"
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
