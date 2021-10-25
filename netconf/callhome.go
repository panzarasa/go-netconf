package netconf

import (
	"net"
    "golang.org/x/crypto/ssh"
)

const (
    CONN_HOST = "0.0.0.0"
    CONN_PORT = "4334"
    CONN_TYPE = "tcp"
)

struct callhomeListener{
	conn net.Conn
	sshConfig *ssh.ClientConfig
}

func (cl *callhomeListener) Initialize(user, pass string) error{
	client_config := ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{ssh.Password(pass)},
        HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            log.Println(hostname)
            log.Println(remote)
            log.Println(key)
            return nil
        },
        Timeout: time.Second * 600,

    }
    sshConfig = &client_config

    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	defer l.Close()
    if err != nil {
        fmt.Println("Error listening:", err.Error())
		return err
    }
    for {
        cl.conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
			return err
        }
        Connection(conn)
    }
}

func (cl *callhomeListener) Connection(){
    sshConn, sshChan, req, _ := ssh.NewClientConn(cl.Conn, "", cl.sshConfig)
    client := ssh.NewClient(sshConn, sshChan, req)
    session,_ := client.NewSession()
    session.RequestSubsystem("netconf")
    var stdoutBuf bytes.Buffer
    session.Stdout = &stdoutBuf
    defer c.Close()
}