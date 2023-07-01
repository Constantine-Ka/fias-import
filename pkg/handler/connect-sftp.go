package handler

import (
	"github.com/pkg/sftp"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"log"
)

//func GetConfig(vp viper.Viper) ssh.Client {
//	user := vp.GetString("SFTP.username")
//	pass := vp.GetString("SFTP.password")
//	remote := vp.GetString("SFTP.host")
//	port := vp.GetString("SFTP.port")
//
//	// get host public key
//	//hostKey := getHostKey(remote)
//
//	config := &ssh.ClientConfig{
//		User: user,
//		Auth: []ssh.AuthMethod{
//			ssh.Password(pass),
//		},
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//		//HostKeyCallback: ssh.FixedHostKey(hostKey),
//	}
//
//	// connect
//	conn, err := ssh.Dial("tcp", remote+":"+port, config)
//	if err != nil {
//		log.Println("ssh")
//		log.Fatal(err)
//	}
//	defer conn.Close()
//
//	return conn
//}

func CreateConnection(vp viper.Viper) (*sftp.Client, *ssh.Client) {
	user := vp.GetString("SFTP.username")
	pass := vp.GetString("SFTP.password")
	remote := vp.GetString("SFTP.host")
	port := vp.GetString("SFTP.port")

	// get host public key
	//hostKey := getHostKey(remote)

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		//HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// connect
	ClientSSH, err := ssh.Dial("tcp", remote+":"+port, config)
	if err != nil {
		log.Println("ssh")
		log.Fatal(err)
	}
	//defer ClientSSH.Close()
	clientSFTP, err := sftp.NewClient(ClientSSH)
	if err != nil {
		log.Println("sftp")
		log.Fatal(err)
	}
	//clientSFTP.Wait()
	//defer clientSFTP.Close()
	return clientSFTP, ClientSSH
}
