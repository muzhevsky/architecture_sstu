package errorHandler

import "fmt"

func HandleError(err error, message string) {
	if err == nil {
		return
	}
	if len(message) == 0 {
		message = err.Error()
	}
	fmt.Println(message)
}
