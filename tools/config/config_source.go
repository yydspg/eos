package config

import (
	"os"

	"github.com/eos/tools/errs"
)

type ConfigSource interface {
	Read() ([]byte, error)
}

type EnvVarSource struct {
	VarName string
}

type FileSystemSource struct {
	FilePath string
}

func (e *EnvVarSource) Read() ([]byte, error) {

	if v, ok := os.LookupEnv(e.VarName); !ok {
		return nil, errs.New("environment variable not set").Wrap()
	} else {
		return []byte(v), nil
	}
}

func (f *FileSystemSource) Read() ([]byte, error) {
	r, err := os.ReadFile(f.FilePath)
	return r, err
}
