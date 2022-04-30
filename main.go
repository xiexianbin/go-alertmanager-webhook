/*
Copyright (c) 2022 xiexianbin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Alarm struct {
	Receiver          string            `json:"receiver"`
	Status            string            `json:"status"`
	Alerts            []Alerts          `json:"alerts"`
	GroupLabels       GroupLabels       `json:"groupLabels"`
	CommonLabels      CommonLabels      `json:"commonLabels"`
	CommonAnnotations CommonAnnotations `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
}
type Labels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Level     string `json:"level"`
}
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
type Alerts struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     time.Time   `json:"startsAt"`
	EndsAt       time.Time   `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
	Fingerprint  string      `json:"fingerprint"`
}
type GroupLabels struct {
	Alertname string `json:"alertname"`
}
type CommonLabels struct {
	Alertname string `json:"alertname"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
	Level     string `json:"level"`
}
type CommonAnnotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%s %s\n", request.Method, request.RequestURI)
		builder := strings.Builder{}
		if request.Method == http.MethodPost {
			fmt.Println("body raw:")
			bs := make([]byte, 1024)
			for {
				n, err := request.Body.Read(bs)
				fmt.Print(string(bs[:n]))
				builder.Write(bs[:n])
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println(err.Error())
				}
			}
			fmt.Println()
		}

		var alarm Alarm
		if err := json.Unmarshal([]byte(builder.String()), &alarm); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("%#v\n", alarm)
		}

		_, _ = writer.Write([]byte(`{"status": 200}`))
	})

	fmt.Println("license on :5001")
	fmt.Println(http.ListenAndServe(":5001", nil))
}
