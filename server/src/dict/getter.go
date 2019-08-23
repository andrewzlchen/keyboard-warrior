package dict

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

// textPayload is a structure that i use to
// unmarshall json data from any api I use for getting words.
// TODO: add more word sources, and add only the key/value
// that I care about for each json payload
type textPayload struct {
	// random text
	Text string `json:"text_out"`
}

// TODO: Figure out how to store this better
const randomTextURL string = "http://www.randomtext.me/api/gibberish/ul-5/5-15"

// GetRandomText makes an api call and gets random
// words and returns them as a list of tokens
func (c *Client) GetRandomText(lowercase bool) ([]string, error) {
	resp, err := c.httpClient.Get(randomTextURL)
	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	payload := textPayload{}
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return []string{}, err
	}

	r := strings.NewReader(payload.Text)
	toks, err := CleanHTMLToTokens(r, lowercase)
	if err != nil {
		return nil, nil
	}

	// convert [][]string to []string
	output := make([]string, 0)
	for i := 0; i < len(toks); i++ {
		for j := 0; j < len(toks[i]); j++ {
			output = append(output, toks[i][j])
		}
	}

	return output, nil
}

// CleanHTMLToTokens converts a HTML string to a 2d slice of tokens
// Expected input: <p>text</p>\r<p>more text</p>
// Expected output: [["text"], ["more", "text"]]
func CleanHTMLToTokens(r io.Reader, lowercase bool) ([][]string, error) {
	output := make([][]string, 0)
	tokenizer := html.NewTokenizer(r)

	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return output, nil
		case html.TextToken:
			// text or \r. Only store text and ignore \r tokens
			t := tokenizer.Token().String()

			// skip if its a whitespace character
			if unicode.IsSpace(rune(t[0])) {
				continue
			}

			toks := processTextToTokens(t, lowercase)
			output = append(output, toks)
		}
	}
}

// processTextToTokens converts a string of text into a list of 1-word tokens
func processTextToTokens(input string, lowercase bool) []string {
	if lowercase {
		// lower case everything
		input = strings.ToLower(input)
	}
	toks := strings.Split(input, " ")
	return toks
}
