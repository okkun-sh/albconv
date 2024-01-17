package cmd_test

import (
	"albconv/cmd"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecompress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		filepath  string
		want      []byte
		wantErr   bool
		errString string
	}{
		{
			"simple",
			"testdata/simple.log.gz",
			[]byte(`http 2018-07-02T22:23:00.186641Z app/my-loadbalancer/50dc6c495c0c9188 192.168.131.39:2817 10.0.0.1:80 0.000 0.001 0.000 200 200 34 366 "GET http://www.example.com:80/ HTTP/1.1" "curl/7.46.0" - - arn:aws:elasticloadbalancing:us-east-2:123456789012:targetgroup/my-targets/73e2d6bc24d8a067 "Root=1-58337262-36d228ad5d99923122bbe354" "-" "-" 0 2018-07-02T22:22:48.364000Z "forward" "-" "-" "10.0.0.1:80" "200" "-" "-"` + "\n"),
			false,
			"",
		},
		{
			"not found file",
			"testdata/not_found.log.gz",
			nil,
			true,
			"open testdata/not_found.log.gz: no such file or directory",
		},
		{
			"not gzip file",
			"testdata/simple.log",
			nil,
			true,
			"file type is not gzip: application/octet-stream",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cmd.Decompress(tt.filepath)
			if err != nil && !tt.wantErr {
				t.Fatalf("Decompress() error = %v", err)
			}
			if err != nil {
				if err.Error() != tt.errString {
					t.Fatalf("Decompress() error = %v, wantErr %v", err, tt.errString)
				}
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatalf("Decompress() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
