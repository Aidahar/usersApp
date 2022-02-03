package ewrap

import "log"

func LogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
