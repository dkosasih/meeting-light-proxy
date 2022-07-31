package v1

// swagger:parameters openhab sendOpenHabCommand
type MeetingStatus struct {
	//  in: body
	//  description: input status json
	//  required: true
	//  example: {"input": {"camera": "inactive", "microphone": "active"}}
	Input struct {
		Camera     string `json:"camera"`
		Microphone string `json:"microphone"`
	} `json:"input"`
}
