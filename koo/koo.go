package koo

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
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

func OokSsh(user string, hostname string, addr string, port uint, command string) {
	// Start new ssh connection with private key.
	//auth := goph.Password(password)
	if hostname == "" {
		hostname = "default"
	}
	//goth.AddKnownHost("alpine",server,,"known_hosts")
	// Start new ssh connection with private key.
	auth, err := goph.Key(".vagrant/machines/"+hostname+"/libvirt/private_key", "")
	if err != nil {
		log.Fatal(err)
	}

	config := &goph.Config{
		User:     user,
		Addr:     addr,
		Port:     port,
		Auth:     auth,
		Callback: VerifyHost,
	}

	client, err := goph.NewConn(config)
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
	authorizedKeysBytes, _ := ioutil.ReadFile("vagrant_private_key")
	//pcert, _, _, _, err := ssh.ParseAuthorizedKey(authorizedKeysBytes)

	host := server + ":22"
	//user = "vagrant"
	//pwd := "vagrant"
	pKey := []byte(authorizedKeysBytes)

	var err error
	var signer ssh.Signer

	signer, err = ssh.ParsePrivateKey(pKey)
	if err != nil {
		fmt.Println(err.Error())
	}
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
