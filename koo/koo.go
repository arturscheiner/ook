package koo

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"time"

	"github.com/janeczku/go-spinner"
	"github.com/k0kubun/go-ansi"
	"github.com/melbahja/goph"
	"github.com/rs/zerolog/log"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/crypto/ssh"
)

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func CheckErr(e error) {
	if e != nil {
		log.Error().Msg(e.Error())
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

func OokSsh(user string, hostname string, addr string, port uint, command string) {
	// Start new ssh connection with private key.
	//auth := goph.Password(password)
	if hostname == "" {
		hostname = "default"
	}
	//goth.AddKnownHost("alpine",server,,"known_hosts")
	// Start new ssh connection with private key.
	auth, err := goph.Key("./.vagrant/machines/"+hostname+"/libvirt/private_key", "")
	CheckErr(err)

	config := &goph.Config{
		User:     user,
		Addr:     addr,
		Port:     port,
		Auth:     auth,
		Callback: VerifyHost,
	}

	client, err := goph.NewConn(config)
	CheckErr(err)

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run(command)
	CheckErr(err)

	// Get your output as []byte.
	fmt.Println(string(out))

}

func ReadFile(fn string, ln int32) (str string, err error) {

	f, err := os.Open(fn)
	CheckErr(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line int
	for scanner.Scan() {
		if line == int(ln) {
			return scanner.Text(), err
		}
		line++
	}
	CheckErr(scanner.Err())

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
	CheckErr(err)

	//io.Copy(io.MultiWriter(f, os.Stdout), stdout)
	//cmd.Wait()

	s.Stop()
}

func SshTest(user string, server string, command string) {
	authorizedKeysBytes, _ := ioutil.ReadFile("vagrant_private_key")
	//pcert, _, _, _, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)

	host := server + ":22"
	//user = "vagrant"
	//pwd := "vagrant"
	pKey := []byte(authorizedKeysBytes)

	var err error
	var signer ssh.Signer

	signer, err = ssh.ParsePrivateKey(pKey)
	CheckErr(err)
	//pukey, _ := ssh.NewPublicKey(signer)
	//goph.AddKnownHost("alpine", server, pukey, "known_hosts")

	//var hostkeyCallback ssh.HostKeyCallback
	//hostkeyCallback, err = knownhosts.New("known_hosts")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	sshConfig := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: VerifyHost,
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	CheckErr(err)

	session, err := client.NewSession()
	CheckErr(err)
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	CheckErr(session.Run(command))

}

func VerifyHost(host string, remote net.Addr, key ssh.PublicKey) error {

	//
	// If you want to connect to new hosts.
	// here your should check new connections public keys
	// if the key not trusted you shuld return an error
	//

	// hostFound: is host in known hosts file.
	// err: error if key not in known hosts file OR host in known hosts file but key changed!
	hostFound, err := goph.CheckKnownHost(host, remote, key, "known_hosts")

	// Host in known hosts but key mismatch!
	// Maybe because of MAN IN THE MIDDLE ATTACK!
	if hostFound && err != nil {

		return err
	}

	// handshake because public key already exists.
	if hostFound && err == nil {

		return nil
	}

	// Ask user to check if he trust the host public key.
	//if askIsHostTrusted(host, key) == false {

	// Make sure to return error on non trusted keys.
	//	return errors.New("you typed no, aborted!")
	//}

	// Add the new host to known hosts file.
	return goph.AddKnownHost(host, remote, key, "known_hosts")
}

func TestOut() {
	// Replace `ls` (and its arguments) with something more interesting
	cmd := exec.Command("vagrant", "halt")

	// Create stdout, stderr streams of type io.Reader
	stdout, err := cmd.StdoutPipe()
	CheckErr(err)
	stderr, err := cmd.StderrPipe()
	CheckErr(err)

	// Start command
	err = cmd.Start()
	CheckErr(err)

	// Don't let main() exit before our command has finished running
	defer cmd.Wait() // Doesn't block

	// Non-blockingly echo command output to terminal
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)

	// I love Go's trivial concurrency :-D
	fmt.Printf("Do other stuff here! No need to wait.\n\n")
}
