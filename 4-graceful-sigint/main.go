//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package sigint

import (
	"os"
	"os/signal"
)

func Start() {
	s := make(chan os.Signal)
	signal.Notify(s, os.Interrupt)
	shutdown := make(chan bool)

	// Create a process
	proc := MockProcess{}

	go func() {
		<-s
		go proc.Stop()
		<-s
		shutdown <- true
	}()

	go func() {
		// Run the process (blocking)
		proc.Run()
		shutdown <- true
	}()
	<-shutdown
}
