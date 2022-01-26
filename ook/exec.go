package ook

import (
	"bufio"
	"ook/koo"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (e *Exec) Run(cmd string) {

	if cmd == "" {
		log.Printf("No command specified!")
		return
	}

	cmdArr := strings.Split(cmd, " ")

	args := []string{}
	if len(cmdArr) > 1 {
		args = cmdArr[1:]
	}

	e.Command = cmdArr[0]
	e.Args = append(e.Args, args...)
	e.Time = zerolog.TimeFormatUnix

	e.execCmd()
}

func (e *Exec) execCmd() {

	command := exec.Command(e.Command, e.Args...)
	command.Env = os.Environ()

	stdout, err := command.StdoutPipe()
	koo.CheckErr(err)

	defer stdout.Close()
	e.Stdout = bufio.NewReader(stdout)

	stderr, err := command.StderrPipe()
	koo.CheckErr(err)

	defer stderr.Close()
	e.Stderr = bufio.NewReader(stderr)

	koo.CheckErr(command.Start())

	go e.handleOutStream()
	go e.handleErrStream()

	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Debug()
				status.ExitStatus()
				return
			}
		}
		return
	}
	return
}
