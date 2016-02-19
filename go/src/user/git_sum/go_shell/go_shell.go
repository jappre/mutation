package goshell

import (
	"bytes"
	"fmt"
	"io"
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
	cmd := exec.Command("/bin/zsh", "-c", command)
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
	// fmt.Printf("%p", *stderr)
	bytesErr, err := readAll(stderr, 1024)
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

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	// fmt.Println(len(buf))
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}
