package netroutine

import "fmt"

func setWorkingData(tokey, value string) string {
	return fmt.Sprintf("set working data at \"%s\" to \"%s\"", tokey, value)
}

func missingWorkingData(fromkey string) string {
	return fmt.Sprintf("unable to find working data at \"%s\"", fromkey)
}

func log(b Runnable, m string, s Status) (string, Status) {
	return fmt.Sprintf("[%s] %s: %s", s.String(), b.kind(), m), s
}
