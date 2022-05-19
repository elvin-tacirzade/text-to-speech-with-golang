package model

type RequestBody struct {
	AudioConfig AudioConfig       `json:"audioConfig"`
	Input       map[string]string `json:"input"`
	Voice       map[string]string `json:"voice"`
}

type ResponseBody struct {
	AudioContent []byte      `json:"audioContent"`
	TimePoints   []string    `json:"timepoints"`
	AudioConfig  AudioConfig `json:"audioConfig"`
}

type AudioConfig struct {
	AudioEncoding    string   `json:"audioEncoding"`
	SpeakingRate     int      `json:"speakingRate"`
	Pitch            int      `json:"pitch"`
	VolumeGainDb     int      `json:"volumeGainDb"`
	SampleRateHertz  int      `json:"sampleRateHertz"`
	EffectsProfileId []string `json:"effectsProfileId"`
}
