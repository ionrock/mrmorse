package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alwindoss/morse"
	"github.com/gin-gonic/gin"
)

// IsMorseCode aims to understand whether a string is morse code or
// not.
func IsMorseCode(msg string) bool {
	validMorseChars := ".- "
	for _, c := range msg {
		if !strings.ContainsRune(validMorseChars, c) {
			return false
		}
	}

	return true
}

type TranslationRequest struct {
	Text string `form:"text"`
}

func handleTranslation(c *gin.Context) {
	var req TranslationRequest
	var translation []byte

	translator := morse.NewHacker()

	// By default we'll translate text to morse code.
	translate := translator.Encode

	c.Bind(&req)

	if IsMorseCode(req.Text) {
		translate = translator.Decode
	}

	// We assume it is text we should translate to morse code
	translation, err := translate(strings.NewReader(req.Text))
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err))
		return
	}

	c.String(http.StatusOK, string(translation))
}

func main() {
	r := gin.Default()

	r.POST("/api/translate", handleTranslation)

	r.Run()
}
