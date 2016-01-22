package goshell

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	// "time"
)

func run() {
	cmd := exec.Command("/bin/sh", "-c", "ping 127.0.0.1")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
}

// GetCommandMessage 用来获取命令行输出的任何内容
func GetCommandMessage(command string) (message string, err error) {
	fmt.Println(command)
	cmd := exec.Command("/bin/sh", "-c", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("StdoutPipe: " + err.Error())
		return err.Error(), err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("StderrPipe: ", err.Error())
		return err.Error(), err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Start: ", err.Error())
		return err.Error(), err
	}

	bytesErr, err := ioutil.ReadAll(stderr)
	if err != nil {
		fmt.Println("ReadAll stderr: ", err.Error())
		return err.Error(), err
	}

	if len(bytesErr) != 0 {
		fmt.Printf("stderr is not nil: %s", bytesErr)
		return string(bytesErr), nil
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ReadAll stdout: ", err.Error())
		return err.Error(), err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait: ", err.Error())
		return err.Error(), err
	}

	// fmt.Printf("stdout: %s", bytes)
	return string(bytes), nil
}
