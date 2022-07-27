package v1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type respondWithReader func(req *http.Request) (*http.Response, error)

func (f respondWithReader) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}
func httpClientWithRoundTripper(statusCode int, response string, wantErr error) *http.Client {
	if wantErr != nil {
		return &http.Client{
			Transport: respondWithReader(func(req *http.Request) (*http.Response, error) {
				return nil, wantErr
			}),
		}
	}

	return &http.Client{
		Transport: respondWithReader(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: statusCode,
				Body:       ioutil.NopCloser(bytes.NewBufferString(response)),
			}, wantErr
		}),
	}
}

func createGinContext(requestJsonPayload string) *gin.Context {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("POST",
		"https://locahost:8443/openhab/command",
		bytes.NewBuffer([]byte(requestJsonPayload)))
	// bytes.NewBuffer([]byte("{\"input\": {\"microphone\":\"inactive\",\"camera\":\"inactive\"}}")))

	return c
}

func TestOpenhabHandler_UpdateOpenHab(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name      string
		fields    *fields
		args      *args
		wantError error
	}{
		{
			name: "Should not have error in gin context when handler is handled well",
			fields: &fields{
				Client: httpClientWithRoundTripper(200, "success message", nil),
			},
			args: &args{
				c: createGinContext("{\"input\": {\"microphone\":\"inactive\",\"camera\":\"inactive\"}}"),
			},
			wantError: nil,
		},
		{
			name: "Should receive parse error when request payload is malformed",
			fields: &fields{
				Client: httpClientWithRoundTripper(200, "success message", nil),
			},
			args: &args{
				c: createGinContext("{\"input\": {\"microphone\":\"inactive\",\"camera\":\"inactive\"} malform"),
			},
			wantError: fmt.Errorf("error parsing meeting status model"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &OpenhabHandler{
				Client: tt.fields.Client,
			}
			h.UpdateOpenHab(tt.args.c)

			if tt.wantError == nil && tt.args.c.Errors != nil {
				t.Errorf("expect: %v; but got: %v", tt.wantError, tt.args.c.Errors[0])
			}

			if tt.wantError != nil &&
				tt.args.c.Errors != nil &&
				!strings.Contains(tt.args.c.Errors[len(tt.args.c.Errors)-1].Error(), tt.wantError.Error()) {
				t.Errorf("expect: %v; but got: %v", tt.wantError, tt.args.c.Errors[0])
			}

		})
	}
}

func Test_sendCommand(t *testing.T) {
	type args struct {
		proxyModel *MeetingStatus
		client     *http.Client
		c          chan error
	}
	tests := []struct {
		name    string
		wantErr error
		args    args
	}{
		{
			name:    "should pass when correct model is passed and server return 200",
			wantErr: nil,
			args: args{
				proxyModel: &MeetingStatus{
					Input: struct {
						Camera     string "json:\"camera\""
						Microphone string "json:\"microphone\""
					}{
						Camera:     "Inactive",
						Microphone: "Inactive",
					},
				},
				client: httpClientWithRoundTripper(200, "success message", nil),
				c:      make(chan error),
			},
		},
		{
			name:    "should error when httpClient responded with non 200",
			wantErr: fmt.Errorf(`Post "https://192.168.0.185:8443/rest/items/MeetingStatus": internal server error`),
			args: args{
				proxyModel: &MeetingStatus{
					Input: struct {
						Camera     string "json:\"camera\""
						Microphone string "json:\"microphone\""
					}{
						Camera:     "Inactive",
						Microphone: "Inactive",
					},
				},
				client: httpClientWithRoundTripper(500, "success message", fmt.Errorf("internal server error")),
				c:      make(chan error),
			},
		},
		{
			name:    "should error when proxy model is null",
			wantErr: fmt.Errorf("Nil proxy model is not allowed"),
			args: args{
				proxyModel: nil,
				client:     httpClientWithRoundTripper(200, "success message", nil),
				c:          make(chan error),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			go func() {
				defer close(tt.args.c)
				sendCommand(tt.args.proxyModel, tt.args.client, tt.args.c)
			}()

			err := <-tt.args.c

			if err != tt.wantErr && err.Error() != tt.wantErr.Error() {
				t.Errorf("expect error: %v; but got: %v", tt.wantErr, err)
			}
		})
	}
}
