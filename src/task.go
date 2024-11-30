package main

//import rl "github.com/gen2brain/raylib-go/raylib"

// gonna work on tasks soon
type TASK struct {
	win       WINDOW
	isRunning bool
	isCleared bool
	pid       int
}

var tasks []TASK
var highestpid uint64 = 0

func Schedular() {
	for i := uint64(0); i < highestpid; i++ {

	}
}
