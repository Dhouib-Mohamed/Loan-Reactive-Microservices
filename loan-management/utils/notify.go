package utils

import (
	"bytes"
	"fmt"
	"net/http"
)

func Notify(bytes_i []byte, message string) {
	// Send Notification
	http.Post("http://notification:8080/notif", "application/json", bytes.NewReader([]byte(message)))
	fmt.Println("Notification Sent :" + message)
	return
}
