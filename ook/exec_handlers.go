package ook

import (
	"bufio"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (e *Exec) handleOutStream() {
	scanner := bufio.NewScanner(e.Stdout)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		e.Stream.Data = scanner.Text()
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Info().Msg(e.Stream.Data)
	}
}

func (e *Exec) handleErrStream() {

	scanner := bufio.NewScanner(e.Stderr)

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		e.Stream.Data = scanner.Text()
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Error().Msg(e.Stream.Data)
	}
}

/* func (e *Exec) handleStream() {

	log.Println(e.Message)
	//log.Println(tools.IsJSONString(message))
	//log.Println(tools.IsJSON(message))

	s := Stream{}
	json.Unmarshal([]byte(e.Message), &s)

	log.Println(s)

	if s.Handler != "" {
		switch {
		case s.Handler == "helm":
			log.Println("Accepted Data Handler -> ", s.Handler)

			// Helm Stream Struct
			hs := helm.Stream(s)

			// Helm Arguments
			ha := helm.Helm{}
			ha.ChartsDir = e.HelmChartsDir
			ha.BinDir = e.HelmBinDir

			ha.CommandsHandlers(hs)

		case s.Handler == "k8s":
			log.Println("Accepted Data Handler")

			ks := k8s.Stream(s)
			k8s.Handlers(ks)

		case s.Handler == "ssh":
			log.Println("Accepted Data Handler")

			// ss := ssh.Stream(s)
			// ssh.Handlers(ss)

		default:
			log.Warn("Unrecognized command")
			e.Conn.Write([]byte("Unrecognized command.\n"))
		}
	} else {
		log.Warn("Invalid input data fields. Should be a JSON encoded data with specific Flee Data fields.")
	}
} */
