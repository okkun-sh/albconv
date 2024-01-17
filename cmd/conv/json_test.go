package conv_test

import (
	"albconv/cmd/conv"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJsonConvert(t *testing.T) {
	tests := []struct {
		name      string
		arg       []byte
		want      func() []byte
		wantErr   bool
		errString string
	}{
		{
			"simple",
			[]byte(`http 2018-07-02T22:23:00.186641Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.000 0.001 0.000 200 200 34 366 "GET http://www.example.com:80/ HTTP/1.1" "curl/7.46.0" - - arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 "Root=1-58337262-36d228ad5d99923122bbe354" "-" "-" 0 2018-07-02T22:22:48.364000Z "forward" "-" "-" "10.0.0.1:80" "200" "-" "-"` + "\n"),
			func() []byte {
				entries := []*conv.ALBLogEntry{
					{
						Type:                   "http",
						Time:                   "2018-07-02T22:23:00.186641Z",
						ELB:                    "app/my-loadbalancer/50dc6c495c0c9188",
						ClientPort:             "192.168.131.39:2817",
						TargetPort:             "10.0.0.1:80",
						RequestProcessingTime:  "0.000",
						TargetProcessingTime:   "0.001",
						ResponseProcessingTime: "0.000",
						ELBStatusCode:          "200",
						TargetStatusCode:       "200",
						ReceivedBytes:          "34",
						SentBytes:              "366",
						Request:                `"GET http://www.example.com:80/ HTTP/1.1"`,
						UserAgent:              `"curl/7.46.0"`,
						SSLCipher:              "-",
						SSLProtocol:            "-",
						TargetGroupARN:         "arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067",
						TraceID:                `"Root=1-58337262-36d228ad5d99923122bbe354"`,
						DomainName:             `"-"`,
						ChosenCertARN:          `"-"`,
						MatchedRulePriority:    "0",
						RequestCreationTime:    "2018-07-02T22:22:48.364000Z",
						ActionsExecuted:        `"forward"`,
						RedirectURL:            `"-"`,
						ErrorReason:            `"-"`,
						TargetPortList:         `"10.0.0.1:80"`,
						TargetStatusCodeList:   `"200"`,
						Classification:         `"-"`,
						ClassificationReason:   `"-"`,
					},
				}
				entriesJSON, _ := json.Marshal(entries)
				return entriesJSON
			},
			false,
			"",
		},
		{
			"invalid number of fields",
			[]byte(`http 2018-07-02T22:23:00.186641Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.000 0.001 0.000 200 200 34 366 "GET http://www.example.com:80/ HTTP/1.1" "curl/7.46.0" - - arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 "Root=1-58337262-36d228ad5d99923122bbe354" "-" "-" 0 2018-07-02T22:22:48.364000Z "forward" "-" "-" "10.0.0.1:80" "-" "-"` + "\n"),
			func() []byte { return nil },
			true,
			"invalid number of fields",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := conv.NewJSONConverter()
			got, err := c.Convert(tt.arg)
			if err != nil && !tt.wantErr {
				t.Fatalf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if err.Error() != tt.errString {
					t.Fatalf("Convert() error = %v, wantErr %v", err, tt.errString)
				}
				return
			}
			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Fatalf("Convert() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
