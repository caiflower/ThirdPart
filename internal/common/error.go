package common

type ThirdError struct {
	Msg string
}

func (e *ThirdError) Error() string {
	if e.Msg == "" {
		return "unKnow error"
	}
	return e.Msg
}
