package cm

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/sftp"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type SFTP interface {
	Upload(lclFile, sftpPath, fileName string) error
	Download(sftpPath, localPath, fileName string) error
}
type klSftp struct {
	cfg *ssh.ClientConfig
}

func NewSFTP() SFTP {
	return &klSftp{
		cfg: config(),
	}
}

func (s *klSftp) Upload(lclFile, sftpPath, fileName string) error {
	client, conn, err := s.getClient()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer client.Close()

	_, err = client.Stat(sftpPath)
	if os.IsNotExist(err) {
		err := client.MkdirAll(sftpPath)
		if err != nil {
			return err
		}
	}
	targetFile := fmt.Sprintf("%v/%v", sftpPath, fileName)
	zap.S().Infof("Start upload file to %v", targetFile)

	dstFile, err := client.Create(targetFile)
	if err != nil {
		zap.S().Errorf("err while create file %v in sftp: %v", targetFile, err)
	}
	defer dstFile.Close()

	srcFile, err := os.Open(lclFile)
	if err != nil {
		zap.S().Errorf("err while create open %v in local: %v", lclFile, err)
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		zap.S().Errorf("err while upload file %v to sftp: %v", lclFile, err)
	}

	return nil
}

func (s *klSftp) Download(sftpPath, localPath, fileName string) error {
	zap.S().Infof("Start download with sftpPath %v localPath %v filename: %v", sftpPath, localPath, fileName)
	client, conn, err := s.getClient()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer client.Close()

	// create destination file
	dstFile, err := os.Create(localPath + fileName)
	zap.S().Infoln("Create file done!")
	// dstFile, err := os.Create("./sftp-download.txt")
	if err != nil {
		zap.S().Errorf("err while create destination file %v in sftp: %v", localPath+fileName, err)
	}
	defer dstFile.Close()

	// open source file
	srcFile, err := client.Open(sftpPath + fileName)
	zap.S().Infoln("Open file done!")
	// srcFile, err := client.Open("./file.txt")
	if err != nil {
		zap.S().Errorf("err while open source file %v in sftp: %v", sftpPath+fileName, err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	zap.S().Infoln("Copy file done!")
	if err != nil {
		zap.S().Errorf("err while copy file to destination in sftp: %v", err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	zap.S().Infoln("Flush memory done!")
	if err != nil {
		zap.S().Errorf("err while flush memory in sftp: %v", err)
	}

	zap.S().Infof("SFTP download file done !!!")
	return nil
}

func (s *klSftp) getClient() (*sftp.Client, *ssh.Client, error) {
	conn, err := ssh.Dial("tcp", os.Getenv("SFTP_HOST"), s.cfg)
	if err != nil {
		zap.S().Errorf("Sftp new Connection error!, SFTP_HOST:%v", os.Getenv("SFTP_HOST"))
		return nil, nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		zap.S().Errorln("Sftp New Client error!")
		return nil, nil, err
	}
	return client, conn, err
}

func config() *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:            os.Getenv("SFTP_USER"),
		Auth:            []ssh.AuthMethod{ssh.Password(os.Getenv("SFTP_PWD"))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}
