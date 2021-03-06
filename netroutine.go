package netroutine

import "fmt"

func reportError(flow string, err error) string {
	return fmt.Sprintf("error during \"%s\": %v", flow, err)
}

func missingSecret(fromkey string) string {
	return fmt.Sprintf("unable to find secret at \"%s\"", fromkey)
}

func reportWrongType(fromkey string) string {
	return fmt.Sprintf("unable to convert the working data at \"%s\" to the required type", fromkey)
}

func setWorkingData(tokey string, value interface{}) string {
	return fmt.Sprintf("set working data at \"%s\" to \"%s\"", tokey, fmt.Sprintf("%v", value))
}

func missingWorkingData(fromkey string) string {
	return fmt.Sprintf("unable to find working data at \"%s\"", fromkey)
}

func log(b Runnable, m string, s Status) (string, Status) {
	return fmt.Sprintf("[%s] %s: %s", s.String(), b.kind(), m), s
}
