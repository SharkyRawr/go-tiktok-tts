package tiktok_tts

import "testing"

func TestTTS(t *testing.T) {
	resp, err := TTS(VoiceEnglishFem1, "Testing")
	if err != nil {
		t.Error(err)
	}
	if len(resp.Data.VStr) <= 0 {
		t.Errorf("VStr is too short: %s", resp.Message)
	}
}
