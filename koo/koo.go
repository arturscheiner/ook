package koo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/janeczku/go-spinner"
	"github.com/k0kubun/go-ansi"
	"github.com/melbahja/goph"
	"github.com/schollz/progressbar/v3"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Bar(len int64, desc string) {
	doneCh := make(chan struct{})

	bar := progressbar.NewOptions(100,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        " ",
			AltSaucerHead: "[yellow]<[reset]",
			SaucerHead:    "[yellow]-[reset]",
			SaucerPadding: "[white]•",
			BarStart:      "[blue]|[reset]",
			BarEnd:        "[blue]|[reset]",
		}),
		progressbar.OptionOnCompletion(func() {
			doneCh <- struct{}{}
		}),
	)

	go func() {
		for i := 0; i < 100; i++ {
			bar.Add(1)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// got notified that progress bar is complete.
	<-doneCh
	fmt.Println("\n ======= progress bar completed ==========\n")
}

func OokSsh(user string, password string, server string, command string) {
	// Start new ssh connection with private key.
	auth := goph.Password(password)

	client, err := goph.NewUnknown(user, server, auth)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run(command)

	if err != nil {
		log.Fatal(err)
	}

	// Get your output as []byte.
	fmt.Println(string(out))
}

func ReadFile(fn string, ln int32) (str string, err error) {

	f, err := os.Open(fn)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line == int(ln) {
			return scanner.Text(), err
		}
		line++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return
}

func Check_up() {
	cmd := exec.Command("vagrant", "ssh", "kv-master-0", "--", "kubectl get nodes -o wide")
	s := spinner.StartNew("This may take some time...")
	//s.SetCharset([]string{"a", "b"})
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//stdout, _ := cmd.StdoutPipe()
	//f, _ := os.Create(".ook/stdout.log")

	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	//io.Copy(io.MultiWriter(f, os.Stdout), stdout)
	//cmd.Wait()

	s.Stop()
}
