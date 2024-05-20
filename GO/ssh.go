package remote

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func RemoteCommandOnce(hostname string, username string, password string, command string) {
	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查
	}

	// 连接远程主机
	client, err := ssh.Dial("tcp", hostname+":22", config)
	if err != nil {
		fmt.Println("Failed to dial: %s", err)
	}
	defer client.Close()

	// 创建一个Session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
	}
	defer session.Close()

	// 执行远程命令
	output, err := session.Output(command)
	if err != nil {
		fmt.Println("Failed to run command: ", err)
	}

	// 打印命令输出
	fmt.Println(string(output))
}

func main() {
	RemoteCommandOnce("127.0.0.1", "root", "xxx", "ls -al")
}
