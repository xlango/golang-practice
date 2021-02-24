package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	for {
		fmt.Println("--- Please press esc ---")
		robotgo.EventHook(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
			fmt.Println("esc")
			robotgo.EventEnd()
		})

		s := robotgo.EventStart()
		<-robotgo.EventProcess(s)

		ok := robotgo.AddEvents("esc")
		if ok {
			fmt.Println("pressed esc...")
		}
	}
}