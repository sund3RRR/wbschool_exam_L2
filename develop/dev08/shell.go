package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

type Shell struct {
	in         *os.File
	out        *os.File
	home       string
	currentDir string
}

func NewShell() (*Shell, error) {
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	home := currentUser.HomeDir

	shell := &Shell{
		in:   os.Stdin,
		out:  os.Stdout,
		home: home,
	}

	shell.currentDir, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	return shell, nil
}

func (shell *Shell) ChangeDirRelative(target string) error {
	splittedCurrentDir := strings.Split(shell.currentDir, "/")
	splittedTargetDir := strings.Split(target, "/")

	for _, str := range splittedTargetDir {
		if str == ".." {
			if len(splittedCurrentDir) == 0 {
				return errors.New("Directory doesn't exist")
			}
			splittedCurrentDir = splittedCurrentDir[:len(splittedCurrentDir)-1]

			currentDir := "/" + strings.Join(splittedCurrentDir, "/")

			err := os.Chdir(currentDir)
			if err != nil {
				return err
			}
		} else {
			splittedCurrentDir = append(splittedCurrentDir, str)
			currentDir := "/" + strings.Join(splittedCurrentDir, "/")
			err := os.Chdir(currentDir)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (shell *Shell) ChangeDirAbsolute(target string) error {
	return os.Chdir(target)
}

func (shell *Shell) ChangeDir(target string) error {
	var err error

	if strings.HasPrefix(target, "/") || strings.HasPrefix(target, "~") {
		err = shell.ChangeDirAbsolute(shell.DeserializePath(target))
	} else {
		err = shell.ChangeDirRelative(target)
	}

	if err != nil {
		return err
	}

	shell.currentDir, err = os.Getwd()
	if err != nil {
		return err
	}

	return nil
}

func (shell *Shell) DeserializePath(path string) string {
	path, found := strings.CutPrefix(path, "~")
	if found {
		return shell.home + path
	}

	return path
}

func (shell *Shell) SerializePath(path string) string {
	path, found := strings.CutPrefix(path, shell.home)
	if found {
		return "~" + path
	}

	return path
}

func (shell *Shell) PrintShellError(err error) {
	fmt.Fprintf(shell.out, "Internal shell error: %v", err)
}

func (shell *Shell) GetPowerWorkingDirectory(args []string) string {
	if len(args) != 0 {
		return fmt.Sprintf("pwd: expected 0 arguments; got %d\n", len(args))
	}
	return shell.currentDir + "\n"
}

func (shell *Shell) Exit() {
	os.Exit(0)
}

func (shell *Shell) Echo(args []string) string {
	return strings.Join(args, " ") + "\n"
}

func (shell *Shell) Kill(args []string) string {
	if len(args) == 0 {
		return "kill: No PID is entered\n"
	}

	for _, arg := range args {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return "kill: PID must be positive integer\n"
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			return fmt.Sprintf("kill: Process %d doesn't exist\n", pid)
		}

		err = process.Kill()
		if err != nil {
			return fmt.Sprintf("kill: Error occured: %v", err)
		}
	}

	return ""
}

func (shell *Shell) Ps(args []string) string {
	if len(args) != 0 {
		return fmt.Sprintf("ps: expected 0 arguments; got %d\n", len(args))
	}

	processes, err := ps.Processes()
	if err != nil {
		return fmt.Sprintf("ps: Error searching process %v\n", err)
	}

	result := make([]string, 0, len(processes)+1)
	result = append(result, "PID\tExecutable")

	for _, pr := range processes {
		line := fmt.Sprintf("%d\t%s", pr.Pid(), pr.Executable())
		result = append(result, line)
	}

	return strings.Join(result, "\n") + "\n"
}

func (shell *Shell) HandleCommand(str string) string {
	splitted := strings.Split(str, " ")
	if len(splitted) < 0 {
		return ""
	}

	command, args := splitted[0], splitted[1:]

	switch command {
	case "cd":
		if len(args) == 0 {
			args = append(args, shell.home)
		}
		err := shell.ChangeDir(args[0])
		if err != nil {
			return fmt.Sprintf("cd: The directory %s doesn't exist\n", args[0])
		}
	case "pwd":
		return shell.GetPowerWorkingDirectory(args)
	case "exit":
		shell.Exit()
	case "echo":
		return shell.Echo(args)
	case "ps":
		return shell.Ps(args)
	case "kill":
		return shell.Kill(args)
	}

	return ""
}
