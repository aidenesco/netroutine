package netroutine

type Status int

const (
	Success Status = iota
	Fail
	Retry
	Error
	Custom
)

func (s Status) String() string {
	return [...]string{"Success", "Fail", "Retry", "Error", "Custom"}[s]
}
