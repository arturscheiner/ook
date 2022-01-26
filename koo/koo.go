package koo

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/janeczku/go-spinner"
	"github.com/k0kubun/go-ansi"
	"github.com/melbahja/goph"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/ssh"
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

func Bar(len int64, desc string, c chan string) {
	bar := progressbar.DefaultBytes(len, desc)
	select {
	case msg := <-c:
		bar.Describe(msg)
		bar.Finish()
	default:
		for i := 0; i < 1000; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
		}
		bar.Finish()

	}
}

func OpBar(fn func()) {
	doneCh := make(chan struct{})

	bar := progressbar.NewOptions(1000,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
		progressbar.OptionOnCompletion(func() {
			doneCh <- struct{}{}
		}),
	)

	go func() {
		for i := 0; i < 1000; i++ {
			bar.Add(1)
			time.Sleep(5 * time.Millisecond)
		}
		fn()
	}()

	// got notified that progress bar is complete.
	<-doneCh
	fmt.Println("\n ======= progress bar completed ==========\n")
}

func OokSsh(user string, password string, server string, command string) {
	// Start new ssh connection with private key.
	//auth := goph.Password(password)

	//goth.AddKnownHost("alpine",server,,"known_hosts")
	// Start new ssh connection with private key.
	auth, err := goph.Key("vagrant_private_key", "abcd")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New(user, server, auth)
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

func SshTest(user string, server string, command string) {
	authorizedKeysBytes, _ := ioutil.ReadFile("vagrant_public_key")
	ssh.ParseAuthorizedKey(authorizedKeysBytes)
	//if err != nil {
	///	log.Printf("Failed to load authorized_keys, err: %v", err)
	//}

	privkeyBytes, _ := ioutil.ReadFile("vagrant_private_key")
	upkey, err := ssh.ParsePrivateKey(privkeyBytes)

	if err != nil {
		log.Printf("Failed to load authorized_keys, err: %v", err)
	}

	usigner, err := ssh.NewSignerFromKey(upkey)
	if err != nil {
		log.Printf("Failed to create new signer, err: %v", err)
	}
	log.Printf("signer: %s", usigner)

	//ucertSigner, err := ssh.NewPublicKey(pcert, usigner)

	if err != nil {
		log.Printf("Failed to create new signer, err: %v", err)
	}

	sshConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(usigner)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", server+":22", sshConfig)

	if err != nil {
		log.Fatalf("Failed to dial, err: %v", err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(command); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	log.Println(b.String())
}
