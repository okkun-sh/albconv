package cmd_test

import (
	"albconv/cmd"
	"albconv/cmd/conv"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		filepaths []string
		want      func() string
		wantErr   bool
		errString string
	}{
		{
			"simple",
			[]string{"testdata/parallel1.log.gz", "testdata/parallel2.log.gz", "testdata/parallel3.log.gz"},
			func() string {
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
					{
						Type:                   "http",
						Time:                   "2019-07-02T22:23:00.186641Z",
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
					{
						Type:                   "http",
						Time:                   "2020-07-02T22:23:00.186641Z",
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
				json, _ := json.Marshal(entries)
				return string(json)
			},
			false,
			"",
		},
		{
			"missing argument",
			[]string{},
			nil,
			true,
			"no files specified",
		},
		{
			"not found file",
			[]string{"testdata/parallel1.log.gz", "testdata/not_found.log.gz", "testdata/parallel3.log.gz"},
			nil,
			true,
			"error processing file testdata/not_found.log.gz: open testdata/not_found.log.gz: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.Run(tt.filepaths)
			if err != nil && !tt.wantErr {
				t.Fatalf("Run() error = %v", err)
			}
			if err != nil {
				if err.Error() != tt.errString {
					t.Fatalf("Run() error = %v, wantErr %v", err, tt.errString)
				}
				return
			}
			if diff := cmp.Diff(got, tt.want()); diff != "" {
				t.Fatalf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
