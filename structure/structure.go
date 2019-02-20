package structure

type BasicInfo struct {
	Server string     `json:"server"`
	Data   RecordInfo `json:"data"`
}

type RecordInfo struct {
	Timestamp string `json:"timestamp"`
	User      int    `json:"user"`
	Qps       int    `json:"qps"`
}
