package wolframalpha

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Answer represents the collective result of a provider query.
type Answer struct {
	Query    string
	Question string
	Answer   string
	Media    []*Media
	Provider string
}

// Media represents a media object in the answer.
type Media struct {
	Type string
	URL  string
}

type WolframProvider struct {
	ApiKey string
}

func NewWolframProvider() *WolframProvider {
	return &WolframProvider{}
}

func (p *WolframProvider) SetApiKey(key string) {
	p.ApiKey = key
}

func (p *WolframProvider) Name() string {
	if p.ApiKey != "" {
		return "Wolfram API"
	}
	return "Wolfram"
}

func (p *WolframProvider) Ask(question string) (string, error) {
	if p.ApiKey == "" {
		return "", errors.New("No API key set")
	}

	u, err := url.Parse("https://api.wolframalpha.com/v2/query")
	if err != nil {
		return "", err
	}
	v := url.Values{}
	v.Set("input", question)
	v.Set("appid", p.ApiKey)
	v.Set("format", "plaintext")
	u.RawQuery = v.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New("Unexpected status: " + resp.Status)
	}

	defer resp.Body.Close()
	dec := xml.NewDecoder(resp.Body)

	result := Queryresult{}
	if err := dec.Decode(&result); err != nil {
		return "", err
	}

	if result.HasError {
		return "", errors.New(result.Error)
	}

	pods := result.Pods
	output := ""
	loopLen := 0
	// hasResult := false

	if len(pods) > 3 {
		loopLen = 2
	} else {
		loopLen = len(pods) - 1
	}

	for i := 0; i <= loopLen; i++ {

		title := pods[i].Title
		textResult := pods[i].SubPods[0].Plaintext
		if title == "Input interpretation" {
			title = "Input"
		}

		if textResult != "" {
			output += fmt.Sprintf("\x02%s:\x02 %s ", title, textResult)
		} else if loopLen < len(pods)-1 {
			loopLen++
		}

	}

	output = strings.Replace(output, "\r", " - ", -1)
	output = strings.Replace(output, "\n", " - ", -1)

	return output, nil
}

// func main() {
// 	wolfram := NewWolframProvider()
// 	wolfram.SetApiKey("5QJJPT-PETXPPHEUY")
// 	_, err := wolfram.Ask("3 mbps * 30 days")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
