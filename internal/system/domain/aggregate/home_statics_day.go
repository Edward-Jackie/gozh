package aggregate

type HomeStaticsDay struct {
	EventDate string `json:"eventDate"`
	Count     int64  `json:"count"`
}

type HomeStaticsDayReply struct {
	Date    []string                       `json:"date"`
	Statics map[string]*HomeStaticsDayInfo `json:"statics"`
}

type HomeStaticsDayInfo struct {
	Date           string `json:"date"`
	VisitCount     int64  `json:"visitCount"`
	VisitUserCount int64  `json:"visitUserCount"`
	PlayCount      int64  `json:"playCount"`
	PlayUserCount  int64  `json:"playUserCount"`
}
