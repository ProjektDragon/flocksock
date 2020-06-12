package reporting

import (
	"encoding/json"
	"fmt"

	"flocksock/log"
)

func FormatMessage(message string) string {
    var formattedMessage string = fmt.Sprintf("flocksock: "+
		"\n"+
		"```"+
		"%s"+
		"```",
		message)

	log.Printf("formatting - Formatted message.\n%s", formattedMessage)
	return formattedMessage
}
