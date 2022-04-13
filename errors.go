package execjs

type RuntimeUnavailableError struct {
	Message string
}

func (e RuntimeUnavailableError) Error() string {
	return e.Message
}

type ProgramError struct {
}

func (e ProgramError) Error() string {
	return "ProgramError"
}
