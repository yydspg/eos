package config 


import (
	"os"
	"path/filepath"
	"github.com/eos/tools/errs"	
)

type PathResolver interface {
	GetDefaultConfigPath() (string,error)
	GetProjectRoot() (string,error)
}
type defaultPathResolver struct {

}

func NewPathResolver() *defaultPathResolver {
	return &defaultPathResolver{}
}

func (d *defaultPathResolver) GetDefaultConfigPath(relativePath string) (string,error) {
	executeablePath,err := os.Executable()
	if err != nil {
		return "",errs.WrapMsg(err,"Executable failed")
	}
	configPath := filepath.Join(filepath.Dir(executeablePath),relativePath)
	return configPath,nil
}

func (d *defaultPathResolver) GetProjectRoot(depth int) (string,error) {
	executePath,err := os.Executable()
	if err != nil {
		return "",errs.WrapMsg(err,"Executeable fail")
	}
	projectRoot := executePath
	for range depth {
		projectRoot = filepath.Dir(projectRoot)
	}
	return projectRoot,nil
}	