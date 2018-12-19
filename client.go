package daemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"github.com/sirupsen/logrus"
)

var goDaemon = flag.Bool("d", false, "run app as a daemon with -d or -d true.")

func Daemon() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *goDaemon {
		//cmd := exec.Command(os.Args[0], flag.Args()[1:]...)
		cmd := exec.Command(os.Args[0])
		cmd.Start()
		logrus.Fatal(fmt.Sprintf("[PID] %d running...\n", cmd.Process.Pid))
		*goDaemon = false
		os.Exit(0)
	}
}
