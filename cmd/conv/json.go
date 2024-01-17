package conv

import (
	"bytes"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
)

type JSONConverter struct {
	reg *regexp.Regexp
}

func NewJSONConverter() *JSONConverter {
	return &JSONConverter{
		reg: regexp.MustCompile(`"[^"]*"|\S+`),
	}
}

func (jc *JSONConverter) Convert(data []byte) ([]byte, error) {
	trimData := bytes.TrimRight(data, "\n")
	var entries []*ALBLogEntry
	for _, l := range strings.Split(string(trimData), "\n") {
		fields := jc.reg.FindAllString(l, -1)
		if len(fields) == 29 {
			entry := &ALBLogEntry{
				Type:                   fields[0],
				Time:                   fields[1],
				ELB:                    fields[2],
				ClientPort:             fields[3],
				TargetPort:             fields[4],
				RequestProcessingTime:  fields[5],
				TargetProcessingTime:   fields[6],
				ResponseProcessingTime: fields[7],
				ELBStatusCode:          fields[8],
				TargetStatusCode:       fields[9],
				ReceivedBytes:          fields[10],
				SentBytes:              fields[11],
				Request:                fields[12],
				UserAgent:              fields[13],
				SSLCipher:              fields[14],
				SSLProtocol:            fields[15],
				TargetGroupARN:         fields[16],
				TraceID:                fields[17],
				DomainName:             fields[18],
				ChosenCertARN:          fields[19],
				MatchedRulePriority:    fields[20],
				RequestCreationTime:    fields[21],
				ActionsExecuted:        fields[22],
				RedirectURL:            fields[23],
				ErrorReason:            fields[24],
				TargetPortList:         fields[25],
				TargetStatusCodeList:   fields[26],
				Classification:         fields[27],
				ClassificationReason:   fields[28],
			}
			entries = append(entries, entry)
		} else {
			return nil, errors.New("invalid number of fields")
		}
	}
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		return nil, err
	}
	return entriesJSON, nil
}
