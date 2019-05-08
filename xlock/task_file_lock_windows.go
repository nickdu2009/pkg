package xlock

type TaskFileLock struct {
}

func NewTaskFileLock(filePath string) (*TaskFileLock, error) {
	return &TaskFileLock{}, nil
}

func (cl *TaskFileLock) Release() error {
	return nil
}

func (cl *TaskFileLock) Lock() error {
	return nil
}

func (cl *TaskFileLock) Unlock() error {
	return nil
}

func (cl *TaskFileLock) ReadFileAll() (string, error) {
	return "", nil
}

func (cl *TaskFileLock) OverwriteFile(b []byte) error {
	return nil
}
