package xlock

import (
	"io"
	"io/ioutil"
	"os"
	"syscall"
)

type TaskFileLock struct {
	filePath string
	file     *os.File
}

func NewTaskFileLock(filePath string) (*TaskFileLock, error) {
	cl := &TaskFileLock{
		filePath: filePath,
	}

	file, err := os.OpenFile(cl.filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	cl.file = file
	return cl, nil
}

func (cl *TaskFileLock) Release() error {
	return cl.file.Close()
}

func (cl *TaskFileLock) Lock() error {
	return syscall.Flock(int(cl.file.Fd()), syscall.LOCK_EX)
}

func (cl *TaskFileLock) Unlock() error {
	return syscall.Flock(int(cl.file.Fd()), syscall.LOCK_UN)
}

func (cl *TaskFileLock) ReadFileAll() (string, error) {
	_, err := cl.file.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}
	fileContentBytes, err := ioutil.ReadAll(cl.file)
	if err != nil {
		return "", err
	}
	return string(fileContentBytes), nil
}

func (cl *TaskFileLock) OverwriteFile(b []byte) error {
	_, err := cl.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = cl.file.Write(b)
	return err
}
