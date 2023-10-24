package util

import "log"

func MustNot(e error) {
	if e != nil {
		log.Fatal(e.Error()) // TODO iikanji
	}
}
