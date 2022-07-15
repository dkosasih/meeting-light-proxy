package openhab

type MeetingStatus struct {
	Input struct {
		Camera     string `json:"camera"`
		Microphone string `json:"microphone"`
	} `json:"input"`
}
