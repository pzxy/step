package common

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"net"
	"runtime"
	"time"
)

func SSH(info *SSHInfo) (map[CmdKey]string, error) {

	sshHost := info.Host
	sshPort := info.Port
	sshUser := info.User
	sshPassword := info.PassWord
	sshType := info.LoginType  //password 或者 key
	sshKeyPath := info.KeyPath //ssh id_rsa.id 路径"

	//创建sshp登陆配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, //ssh 连接time out 时间一秒钟, 如果ssh验证错误 会在一秒内返回
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //这个可以， 但是不够安全
		//HostKeyCallback: hostKeyCallBackFunc(h.Host),
	}
	if sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	//dial 获取ssh client
	addr := fmt.Sprintf("%s:%s", sshHost, sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("创建ssh client 失败:%v", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		return nil, fmt.Errorf("创建ssh session 失败:%v", err)
	}
	defer session.Close()
	//执行远程命令
	var ret = make(map[CmdKey]string, 0)
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

func GetMacAddrBySSH(info *SSHInfo) (string, error) {
	cmd := make(map[CmdKey]string, 0)
	switch runtime.GOOS {
	case Linux:
		cmd[HwAddr] = LinuxMacShell
	case Darwin:
		cmd[HwAddr] = DarwinMacShell
	default:
		return "", fmt.Errorf("%s操作系统的 GetMacAddr方法 尚未实现。", runtime.GOOS)
	}
	order := []CmdKey{HwAddr}
	info.Cmd = cmd
	info.CmdOrder = order

	hwMap, err := SSH(info)
	if err != nil {
		return "", err
	}
	return hwMap[HwAddr], nil
}

func GetMacAddrByLocalhost() (string, error) {

	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		//mac := inter.HardwareAddr //获取本机MAC地址
		hwAddr := inter.HardwareAddr.String()
		if hwAddr == "" {
			continue
		}
		return hwAddr, nil
	}

	return "", fmt.Errorf("not found mac addr")
}
