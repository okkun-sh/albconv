package conv

type ALBLogEntry struct {
	Type                   string `json:"type"`
	Time                   string `json:"time"`
	ELB                    string `json:"elb"`
	ClientPort             string `json:"client_port"`
	TargetPort             string `json:"target_port"`
	RequestProcessingTime  string `json:"request_processing_time"`
	TargetProcessingTime   string `json:"target_processing_time"`
	ResponseProcessingTime string `json:"response_processing_time"`
	ELBStatusCode          string `json:"elb_status_code"`
	TargetStatusCode       string `json:"target_status_code"`
	ReceivedBytes          string `json:"received_bytes"`
	SentBytes              string `json:"sent_bytes"`
	Request                string `json:"request"`
	UserAgent              string `json:"user_agent"`
	SSLCipher              string `json:"ssl_cipher"`
	SSLProtocol            string `json:"ssl_protocol"`
	TargetGroupARN         string `json:"target_group_arn"`
	TraceID                string `json:"trace_id"`
	DomainName             string `json:"domain_name"`
	ChosenCertARN          string `json:"chosen_cert_arn"`
	MatchedRulePriority    string `json:"matched_rule_priority"`
	RequestCreationTime    string `json:"request_creation_time"`
	ActionsExecuted        string `json:"actions_executed"`
	RedirectURL            string `json:"redirect_url"`
	ErrorReason            string `json:"error_reason"`
	TargetPortList         string `json:"target_port_list"`
	TargetStatusCodeList   string `json:"target_status_code_list"`
	Classification         string `json:"classification"`
	ClassificationReason   string `json:"classification_reason"`
}

type Converter interface {
	Convert(t string, data []byte) ([]byte, error)
}
