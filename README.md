## TikTok Text-To-Speech module for Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/SharkyRawr/go-tiktok-tts.svg)](https://pkg.go.dev/github.com/SharkyRawr/go-tiktok-tts)

### Example:
```go
    tts, err := tiktok_tts.TTS(tiktok_tts.VoiceEnglishFem1, "Hello World")
    if err != nil { /* ... */ }

    buf, err := base64.StdEncoding.DecodeString(tts.Data.VStr)
    if err != nil { /* ... */ }
    // buf now contains MP3 file data which can be written to a file
```
