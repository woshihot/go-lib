package os

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(name string, args ...string) (string, error) {

	return cmd(name, true, args...)
}

func RunCmdNoWait(name string, args ...string) (string, error) {
	return cmd(name, false, args...)
}

func cmd(name string, wait bool, arg ...string) (string, error) {

	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("[runCmd-error] %s\n", err)
		return "", err
	}

	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if wait {
		cmd.Wait()
	}

	return string(out[:]), nil
}

func RunWaitCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("[runCmd-error] %s\n", err)
		return "", err
	}

	out, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	cmd.Wait()
	return string(out[:]), nil

}

func RunCmdWaitResultByKey(resultKey, name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("[runCmd-error] %s\n", err)
		return "", err
	}

	reader := bufio.NewReader(stdout)
	firstLine := ""
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			log.Printf("end of file or err = %v", err2)
			break
		}
		if strings.HasPrefix(line, resultKey) {
			firstLine = strings.TrimPrefix(line, resultKey)
			log.Printf("firstLine = %s\n", firstLine)
			break
		}

		//if firstLine != "" {
		//	break
		//}
	}
	cmd.Wait()
	return strings.TrimSpace(firstLine), nil
	//out, err := ioutil.ReadAll(stdout)
	//
	//err = cmd.Wait()
	//if err != nil {
	//	return "", err
	//}
	//return string(out[:]), nil
}
