package logger

import "time"

// LogErr prints content of error to console
// with additional information.
// Returns true, if error exists
func LogErr(err error) bool {
	if err != nil {
		message := time.Now().Format(time.RFC822) + ";    ERROR:\n" + err.Error()
		println(message)
		return true
	}
	return false
}

func LogData(data string) {
	message := time.Now().Format(time.RFC822) + ";    DATA:\n" + data
	println(message)
}
