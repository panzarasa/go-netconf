package netconf

import (
	"net"
    "golang.org/x/crypto/ssh"
    "fmt"
    "time"
    "bytes"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "4334"
    CONN_TYPE = "tcp"
)

type CallhomeListener struct{
	conn net.Conn
	sshConfig *ssh.ClientConfig
}

func (cl *CallhomeListener) Initialize(user, pass string) error{
	client_config := ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{ssh.Password(pass)},
        HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            fmt.Println(hostname)
            fmt.Println(remote)
            fmt.Println(key)
            return nil
        },
        Timeout: time.Second * 600,

    }
    cl.sshConfig = &client_config

    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	defer l.Close()
    if err != nil {
        fmt.Println("Error listening:", err.Error())
		return err
    }
    for {
        cl.conn, err = l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
			return err
        }
        cl.Connection()
    }
}

func (cl *CallhomeListener) Connection(){
    sshConn, sshChan, req, _ := ssh.NewClientConn(cl.conn, "", cl.sshConfig)
    client := ssh.NewClient(sshConn, sshChan, req)
    session,_ := client.NewSession()
    session.RequestSubsystem("netconf")
    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
}