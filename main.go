package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"typer/color"
)

func main() {
	startTyper()
}

func startTyper() {
	var phase string = getQuote()
	var letters []string

	var done bool = true

	fmt.Println("Press any letter to start!")
	fmt.Println(phase)

	for done {
		keyboard.Listen(func(key keys.Key) (stop bool, err error) {
			switch key.Code {
			case keys.CtrlC:
				done = stopTyper()
			case keys.Backspace:
				if len(letters) > 0 {
					letters = letters[:len(letters)-1]
				}
			case keys.Space:
				letters = append(letters, " ")
			case keys.Enter:
				if strings.Join(letters, "") == phase {
					phase, letters = getQuote(), []string{}
				}
			default:
				letters = appendKey(letters, phase, key.String())
			}

			clearScreen()
			drawScreen(phase, letters)

			return true, nil
		})
	}
}

func stopTyper() bool {
	fmt.Printf("\n%s", color.RESET)
	return false
}

func appendKey(letters []string, phase string, key string) []string {
	var lastPhaseChar string = string(phase[len(letters)])

	if len(letters) >= len(phase) {
		letters = append(letters, color.Colorize(key, color.RED))
		return letters
	}
	if lastPhaseChar != key {
		letters = append(letters, color.Colorize(key, color.RED))
		return letters
	}
	if lastPhaseChar == key {
		letters = append(letters, key)
		return letters
	}

	letters = append(letters, key)
	return letters
}

func drawScreen(phase string, letters []string) {
	fmt.Println(phase)
	fmt.Print(strings.Join(letters, ""))
}

func clearScreen() {
	fmt.Print("\033[2J\033[1;1H")
}

func getQuote() string {
	resp, err := http.Get("https://recite.onrender.com/random/quote")
	if err != nil {
		return "error: " + err.Error()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "error: " + err.Error()
	}

	var quoteResponse QuoteResponse
	err = json.Unmarshal(body, &quoteResponse)
	if err != nil {
		return "error: " + err.Error()
	}

	return quoteResponse.Data.Quote
}

type QuoteData struct {
	Quote     string    `json:"quote"`
	Book      string    `json:"book"`
	Author    string    `json:"author"`
	Length    int       `json:"length"`
	Words     int       `json:"words"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QuoteResponse struct {
	ID   string    `json:"id"`
	Data QuoteData `json:"data"`
}
