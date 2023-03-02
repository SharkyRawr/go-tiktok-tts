// Golang module to generate text-to-speech using the TikTok API.
package tiktok_tts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Voice string

const (
	// US English Female 1
	VoiceEnglishFem1 = "en_us_001"
	// US English Female 2
	VoiceEnglishFem2 = "en_us_002"
	// US English Male 1
	VoiceEnglishMal1 = "en_us_006"
	// US English Male 2
	VoiceEnglishMal2 = "en_us_007"
	// US English Male 3
	VoiceEnglishMal3 = "en_us_009"
	// US English Male 4
	VoiceEnglishMal4 = "en_us_010"

	// Australian English Female
	VoiceEnglishFemAU = "en_au_001"
	// Australian English Male
	VoiceEnglishMalAU = "en_au_002"

	// UK English Male 1
	VoiceEnglishMalGB = "en_uk_001"
	// UK English Male 2
	VoiceEnglishMalGB2 = "en_uk_003"

	// French Male 1
	VoiceFrenchMal1 = "fr_001"
	// French Male 2
	VoiceFrenchMal2 = "fr_002"

	// German Female
	VoiceGermanFem = "de_001"
	// German Male
	VoiceGermanMal = "de_002"

	// Spanish Male
	VoiceEspMal = "es_002"

	// Spanish (Mexican) Male
	VoiceSpaMal = "es_mx_002"

	// Brazilian Female 1
	VoiceBraFem1 = "br_001"
	// Brazilian Female 2
	VoiceBraFem2 = "br_003"
	// Brazilian Female 3
	VoiceBraFem3 = "br_004"
	// Brazilian Male
	VoiceBraMal = "br_005"

	// Idonesian Male
	VoiceIdoFem = "id_001"

	// Japanese Female 1
	VoiceJpnFem1 = "jp_001"
	// Japanese Female 2
	VoiceJpnFem2 = "jp_003"
	// Japanese Female 3
	VoiceJpnFem3 = "jp_005"
	// Japanese Male
	VoiceJpnMal = "jp_006"

	// Korean Male 1
	VoiceKorMal1 = "kr_002"
	// Korean Female
	VoiceKorFem = "kr_003"
	// Korean Male 2
	VoiceKorMal2 = "kr_004"

	// Disney Ghostface
	VoiceGhostface = "en_us_ghostface"
	// Disney Chewbacca
	VoiceChewbacca = "en_us_chewbacca"
	// Disney C3P0
	VoiceC3PO = "en_us_c3po"
	// Disney Stitch
	VoiceStitch = "en_us_stitch"
	// Disney Stormtrooper
	VoiceStormtrooper = "en_us_stormtrooper"
	// Disney Rocket
	VoiceRocket = "en_us_rocket"
)

const api_url = "https://api16-normal-useast5.us.tiktokv.com/media/api/text/speech/invoke/?text_speaker=%s&req_text=%s&speaker_map_type=0&aid=1233"

// Text-to-Speech Response JSON struct
type TTSResponse struct {
	Data       Data   `json:"data"`
	Message    string `json:"message"`
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// Text-to-speech response Data struct
// Vstr is the voice mp3 data in base64
type Data struct {
	SKey     string `json:"s_key"`
	VStr     string `json:"v_str"`
	Duration string `json:"duration"`
}

// Generator Text-to-Speech MP3 data from string text with voice voice.
//
// Example:
//
//	tts, err := tiktok_tts.TTS(tiktok_tts.VoiceEnglishFem1, "Hello World")
//	if err != nil { /* ... */ }
//	buf, err := base64.StdEncoding.DecodeString(tts.Data.VStr)
//	if err != nil { /* ... */ }
//	// buf now contains MP3 file data which can be written to a file
func TTS(voice Voice, text string, sessionid string) (TTSResponse, error) {
	text = strings.ReplaceAll(text, "+", "plus")
	text = strings.ReplaceAll(text, " ", "+")
	text = strings.ReplaceAll(text, "&", "and")

	requestUrl := fmt.Sprintf(api_url, voice, text)
	jar, err := cookiejar.New(nil)
	if err != nil {
		return TTSResponse{}, err
	}

	jar.SetCookies(&url.URL{
		Scheme: "https",
		Host:   "api16-normal-useast5.us.tiktokv.com",
	}, []*http.Cookie{{
		Name:  "sessionid",
		Value: sessionid,
	},
	})

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Post(requestUrl, "text/plain", nil)
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
