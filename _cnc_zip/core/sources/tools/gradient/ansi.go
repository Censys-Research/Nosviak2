package gradient 

import (
	"regexp"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func strip(str string) string {
	return re.ReplaceAllString(str, "")
}

func contains(str string) bool {
	return re.MatchString(str)
}

func get(str string) string {
	return re.FindString(str)
}