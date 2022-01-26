package koo

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Execute(cmd string) (err error) {
	if cmd == "" {
		return errors.New("No command provided")
	}

	cmdArr := strings.Split(cmd, " ")
	name := cmdArr[0]

	args := []string{}
	if len(cmdArr) > 1 {
		args = cmdArr[1:]
	}

	command := exec.Command(name, args...)
	command.Env = os.Environ()

	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Error()
		return err
	}
	defer stdout.Close()
	stdoutReader := bufio.NewReader(stdout)

	stderr, err := command.StderrPipe()
	if err != nil {
		log.Error()
		return err
	}
	defer stderr.Close()
	stderrReader := bufio.NewReader(stderr)

	if err := command.Start(); err != nil {
		log.Error()
		return err
	}

	go HandleReader(stdoutReader)
	go HandleReader(stderrReader)

	if err := command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				log.Debug()
				status.ExitStatus()
				return err
			}
		}
		return err
	}
	return nil
}

func HandleReader(reader *bufio.Reader) {
	scanner := bufio.NewScanner(reader)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		message := scanner.Text()
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Print(message)
		//fmt.Println(message)
	}
}
