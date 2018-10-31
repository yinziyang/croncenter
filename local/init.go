package local

import (
	"os/user"
	"path/filepath"
)

const (
	PUBLIC_KEY_FILENAME  = "id_rsa.pub"
	PRIVATE_KEY_FILENAME = "id_rsa"
	SSH_DIR_NAME         = ".ssh"
)

var (
	UserName       string
	HomeDir        string
	SSHdir         string
	PrivateKeyFile string
	PublicKeyFile  string
)

func init() {
	if u, err := user.Current(); err != nil {
		panic(err)
	} else {
		UserName = u.Username
		HomeDir = u.HomeDir
	}

	SSHdir = filepath.Join(HomeDir, SSH_DIR_NAME)

	PrivateKeyFile = filepath.Join(SSHdir, PRIVATE_KEY_FILENAME)
	PublicKeyFile = filepath.Join(SSHdir, PUBLIC_KEY_FILENAME)
}
