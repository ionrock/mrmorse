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
	validMorseChars := ".- /|"
	for _, c := range msg {
		if !strings.ContainsRune(validMorseChars, c) {
			return false
		}
	}

	return true
}

// Translate encodes or decodes the text to and from morse code.
func Translate(input string) (string, error) {
	var translation []byte

	translator := morse.NewHacker()

	// By default we'll translate text to morse code.
	translate := translator.Encode

	if IsMorseCode(input) {
		translate = translator.Decode
	}

	// We assume it is text we should translate to morse code
	translation, err := translate(strings.NewReader(input))
	if err != nil {
		return input, err
	}

	return string(translation), nil
}

type TranslationRequest struct {
	Text string `form:"text"`
}

func handleTranslation(c *gin.Context) {
	var req TranslationRequest

	c.Bind(&req)

	translation, err := Translate(req.Text)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err))
		return
	}

	c.String(http.StatusOK, string(translation))
}

func main() {
	r := gin.Default()

	r.POST("/api/translate", handleTranslation)
	r.GET("/api/translate", handleTranslation)

	r.Run()
}
