package util

import (
	"example.com/m/v2/common"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func SSH(info *common.SSHInfo) (map[string]string, error) {
	sshClient, err := newSshClient(info)
	if err != nil {
		return nil, fmt.Errorf("创建ssh client 失败:%v", err)
	}
	defer sshClient.Close()

	session, err := sshClient.NewSession()
	if err != nil {
		return nil, fmt.Errorf("创建ssh session 失败:%v", err)
	}
	defer session.Close()
	//执行远程命令
	var ret = make(map[string]string, 0)
	for _, k := range info.CmdOrder {
		v, ok := info.Cmd[k]
		if !ok {
			return nil, fmt.Errorf("cmd %v can't find", k)
		}
		output, err := session.CombinedOutput(v)
		if err != nil {
			return nil, err
		}
		ret[k] = string(output)
	}
	return ret, nil
}

func newSshClient(info *common.SSHInfo) (*ssh.Client, error) {
	sshHost := info.Host
	sshPort := info.Port
	sshUser := info.User
	sshType := info.LoginType //password 或者 key

	//创建ssh登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second * 3, //ssh 连接time out 时间3秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == common.PassWordType {
		config.Auth = []ssh.AuthMethod{ssh.Password(info.LoginData)} //password 或者 key
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(info.LoginData)} //ssh id_rsa.id 路径"
	}

	addr := fmt.Sprintf("%s:%s", sshHost, sshPort)
	return ssh.Dial("tcp", addr, config)
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func ScpFile(src, dst string, info *common.SSHInfo) error {
	var err error
	sshClient, err := newSshClient(info)
	if err != nil {
		return err
	}
	defer sshClient.Close()
	client, err := scp.NewClientBySSH(sshClient)

	if err != nil {
		return err
	}
	err = client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return err
	}
	defer client.Close()

	f, _ := os.Open(src)

	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, dst, "0655")

	if err != nil {
		fmt.Println("Error while copying file ", err)
		return err
	}
	return nil
}

func DoSsh(cmd string, info *common.SSHInfo) (string, error) {
	if cmd == "" || info == nil {
		return "", fmt.Errorf("Invalid parameter")
	}
	o := "o"
	c := make(map[string]string, 0)
	c[o] = cmd
	info.CmdOrder = []string{o}
	info.Cmd = c
	m, err := SSH(info)
	return m[o], err
}
