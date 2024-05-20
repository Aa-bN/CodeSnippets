package remote

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func RemoteShell(hostname string, username string, password string) {
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
		fmt.Println("Failed to dial: ", err)
	}
	defer client.Close()

	// 创建一个Session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
	}
	defer session.Close()

	// 设置终端模式
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // 禁用回显
		ssh.TTY_OP_ISPEED: 14400, // 输入速度 = 14.4 kbaud
		ssh.TTY_OP_OSPEED: 14400, // 输出速度 = 14.4 kbaud
	}

	// 请求交互式终端
	if err := session.RequestPty("linux", 80, 40, modes); err != nil {
		fmt.Println("Failed to request pty: ", err)
	}

	//设置输入输出
	session.Stdout = os.Stdout
	session.Stdin = os.Stdin
	session.Stderr = os.Stderr

	if err := session.Shell(); err != nil {
		fmt.Println("Failed to start shell: ", err)
	}

	err = session.Wait()
	if err != nil {
		fmt.Println("Failed to run: " + err.Error())
	}

	fmt.Println("Remote session finished")
}
