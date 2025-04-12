package status

// Status represents transcription status
type Status int

const (
	// Uploaded value
	Uploaded Status = iota + 1
	// SplitChannels value
	SplitChannels
	// AudioConvert value
	AudioConvert
	// Diarization value
	Diarization
	// Transcription value
	Transcription
	// Rescore status
	Rescore
	// Whisper status
	Whisper
	// ResultMake status
	ResultMake
	// Res2Eaf status
	Res2Eaf
	// JoinResults status
	JoinResults
	// Completed status
	Completed
)

var (
	statusName = map[Status]string{Uploaded: "UPLOADED", Completed: "COMPLETED",
		SplitChannels: "SplitChannels", AudioConvert: "AudioConvert", Diarization: "Diarization",
		Transcription: "Transcription", Rescore: "Rescore",
		ResultMake: "ResultMake", JoinResults: "JoinResults", Whisper: "Whisper", Res2Eaf: "Res2Eaf"}
	nameStatus = map[string]Status{"UPLOADED": Uploaded, "COMPLETED": Completed,
		"SplitChannels": SplitChannels,
		"AudioConvert":  AudioConvert, "Diarization": Diarization,
		"Transcription": Transcription, "Rescore": Rescore,
		"ResultMake": ResultMake, "JoinResults": JoinResults, "Whisper": Whisper, "Res2Eaf": Res2Eaf}
)

// Name return status as string
func Name(st Status) string {
	return statusName[st]
}

// From converts string to Status
func From(st string) Status {
	return nameStatus[st]
}

// Min selects min status of the two
func Min(st1, st2 Status) Status {
	if st1 < st2 {
		return st1
	}
	return st2
}
