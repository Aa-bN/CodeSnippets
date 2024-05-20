package remote

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func RemoteUploadFile(hostname string, username string, password string, local_file_path string, remote_file_path string) {
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

	// 创建SFTP客户端
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("Failed to create SFTP client: ", err)
	}
	defer sftpClient.Close()

	// 上传文件
	localFile, err := os.Open(local_file_path)
	if err != nil {
		fmt.Println("Failed to open local file: ", err)
	}
	defer localFile.Close()

	remoteFile, err := sftpClient.Create(remote_file_path)
	if err != nil {
		fmt.Println("Failed to create remote file: ", err)
	}
	defer remoteFile.Close()

	_, err = io.Copy(remoteFile, localFile)
	if err != nil {
		fmt.Println("Failed to upload file: ", err)
	}

	fmt.Println("File uploaded successfully")
}

func RemoteDownloadFile(hostname string, username string, password string, remote_file_path string, local_file_path string) {
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

	// 创建SFTP客户端
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("Failed to create SFTP client: ", err)
	}
	defer sftpClient.Close()

	// 下载文件
	remoteFile, err := sftpClient.Open(remote_file_path)
	if err != nil {
		fmt.Println("Failed to open remote file: ", err)
	}
	defer remoteFile.Close()

	localFile, err := os.Create(local_file_path)
	if err != nil {
		fmt.Println("Failed to create local file: ", err)
	}
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		fmt.Println("Failed to download file: ", err)
	}

	fmt.Println("File downloaded successfully")
}
