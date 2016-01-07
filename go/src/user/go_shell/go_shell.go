package main

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

func main() {
	// go run()
	// time.Sleep(1e9)
	//
	// cmd := exec.Command("/bin/sh", "-c", "cd ~;l")
	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	fmt.Println("StdoutPipe: " + err.Error())
	// 	return
	// }
	//
	// stderr, err := cmd.StderrPipe()
	// if err != nil {
	// 	fmt.Println("StderrPipe: ", err.Error())
	// 	return
	// }
	//
	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("Start: ", err.Error())
	// 	return
	// }
	//
	// bytesErr, err := ioutil.ReadAll(stderr)
	// if err != nil {
	// 	fmt.Println("ReadAll stderr: ", err.Error())
	// 	return
	// }
	//
	// if len(bytesErr) != 0 {
	// 	fmt.Printf("stderr is not nil: %s", bytesErr)
	// 	return
	// }
	//
	// bytes, err := ioutil.ReadAll(stdout)
	// if err != nil {
	// 	fmt.Println("ReadAll stdout: ", err.Error())
	// 	return
	// }
	//
	// if err := cmd.Wait(); err != nil {
	// 	fmt.Println("Wait: ", err.Error())
	// 	return
	// }
	//
	// fmt.Printf("stdout: %s", bytes)
	getCommandMessage("cd ~;ls")
}

func getCommandMessage(command string) (message string, err error) {
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

	fmt.Printf("stdout: %s", bytes)
	return string(bytes), nil
}
