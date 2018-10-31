package sshtools

import (
	"bytes"
	"croncenter/local"
	"io/ioutil"
	"os"
	"path"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func CreateTrust(remoteServer, userName, passwd string) (err error) {
	// 尝试本机生成密钥
	if err = local.GenSSHkey(); err != nil {
		return
	}

	var sshClient *ssh.Client
	// 如果已经建立互信, 直接返回
	if sshClient, err = GetSSHClient(remoteServer, userName); err == nil {
		return
	}

	// 尝试通过密码建立ssh连接
	if sshClient, err = GetSSHClientByPasswd(remoteServer, userName, passwd); err != nil {
		return
	}
	defer sshClient.Close()

	var session *ssh.Session
	if session, err = sshClient.NewSession(); err != nil {
		return
	}
	defer session.Close()

	// 获取远端.ssh路径
	var remoteHomeDir []byte
	if remoteHomeDir, err = session.Output("echo $HOME"); err != nil {
		return
	}
	remoteHomeDir = bytes.TrimSpace(remoteHomeDir)
	remoteSshDir := path.Join(string(remoteHomeDir), ".ssh")
	remoteAuthorizedKeysFile := path.Join(remoteSshDir, "authorized_keys")

	// 将本机publickey同步到远端服务器建立互信
	var publicKey []byte
	if publicKey, err = local.GetPublicKey(); err != nil {
		return
	}

	var client *sftp.Client
	if client, err = sftp.NewClient(sshClient); err != nil {
		return
	}
	defer client.Close()
	if err = client.MkdirAll(remoteSshDir); err != nil {
		return
	}
	if err = client.Chmod(remoteSshDir, 0700); err != nil {
		return
	}

	var f *sftp.File
	if f, err = client.OpenFile(remoteAuthorizedKeysFile, os.O_CREATE|os.O_APPEND|os.O_RDWR); err != nil {
		return
	}
	defer f.Close()
	if _, err = ioutil.ReadAll(f); err != nil {
		return
	}
	if _, err = f.Write(append(publicKey, byte('\n'))); err != nil {
		return
	}
	if err = client.Chmod(remoteAuthorizedKeysFile, 0644); err != nil {
		return
	}

	return
}
