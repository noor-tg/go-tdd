package main

import (
	"alnoor/gotdd/mocking"
	"os"
)

// for dep_inject
// func main() {
// 	dep_inject.Terminal()
// 	dep_inject.Web()
// }

func main() {
	mocking.Countdown(os.Stdout)
}
