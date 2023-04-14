package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_processUrls(t *testing.T) {
	buff := bytes.Buffer{}
	log.SetOutput(&buff)
	defer log.SetOutput(os.Stdout)

	tests := []struct {
		name        string
		workerCount int
		urls        []string
		expected    []string
	}{
		{
			name:        "single url",
			workerCount: 1,
			urls:        []string{"/about"},
			expected:    []string{"/about 7023e8e14be3181c54f46bbce3efc5b8"},
		},
		{
			name:        "number of urls more then workers",
			workerCount: 2,
			urls:        []string{"/about", "/users", "/contacts"},
			expected: []string{
				"/about 7023e8e14be3181c54f46bbce3efc5b8",
				"/users d4e6a2efb256af652053607257fe1d14",
				"/contacts f5ce3a2238cea1472d49974490e0126b",
			},
		},
		{
			name:        "number of workers more then urls",
			workerCount: 5,
			urls:        []string{"/about", "/users", "/contacts"},
			expected: []string{
				"/about 7023e8e14be3181c54f46bbce3efc5b8",
				"/users d4e6a2efb256af652053607257fe1d14",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer buff.Reset()

			// create a new test server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte(fmt.Sprintf("Response from %s url", r.URL.Path)))
			}))
			defer ts.Close()

			// Replace test URLs with local server URL
			for i := range tt.urls {
				tt.urls[i] = ts.URL + tt.urls[i]
			}

			processUrls(tt.workerCount, tt.urls)

			for _, expected := range tt.expected {
				expectedOutput := ts.URL + expected

				if !strings.Contains(buff.String(), expectedOutput) {
					t.Errorf("Unexpected result: expected %q, got %q", expectedOutput, buff.String())
				}
			}
		})
	}
}
