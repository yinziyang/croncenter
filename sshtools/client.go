package sshtools

import (
	"croncenter/local"

	"golang.org/x/crypto/ssh"
)

func getSSHClientByAuth(remoteServer string, userName string, auth []ssh.AuthMethod) (client *ssh.Client, err error) {
	client, err = ssh.Dial("tcp", remoteServer+":22",
		&ssh.ClientConfig{
			User:            userName,
			Auth:            auth,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
	return
}

func GetSSHClient(remoteServer, userName string) (client *ssh.Client, err error) {
	var auth []ssh.AuthMethod
	if auth, err = local.GetSSHAuth(); err != nil {
		return
	}
	if client, err = getSSHClientByAuth(remoteServer, userName, auth); err != nil {
		return
	}
	return
}

func GetSSHClientByPasswd(remoteServer, userName, passwd string) (client *ssh.Client, err error) {
	auth := []ssh.AuthMethod{
		ssh.PasswordCallback(func() (string, error) {
			return passwd, nil
		}),
	}
	if client, err = getSSHClientByAuth(remoteServer, userName, auth); err != nil {
		return
	}
	return
}
