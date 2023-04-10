package callback

type AliyunMediaCallback[T any] struct {
	EventType   string `json:"EventType"`
	UserId      int64  `json:"UserId"`
	EventTime   string `json:"EventTime"`
	MessageBody T      `json:"MessageBody"`
}

type AliyunTranscodingMediaStruct struct {
	ParentJobId   string                       `json:"ParentJobId"`
	Jobs          []AliyunTranscodingMediaJobs `json:"Jobs"`
	TriggerSource string                       `json:"TriggerSource"`
	Name          string                       `json:"Name"`
	FinishTime    string                       `json:"FinishTime"`
	UserData      string                       `json:"UserData"`
	Status        string                       `json:"Status"`
}

type AliyunTranscodingMediaJobs struct {
	Status    string `json:"Status"`
	JobId     string `json:"JobId"`
	OutputUrl string `json:"OutputUrl"`
}
