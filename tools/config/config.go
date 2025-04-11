package config

import (
	"os"
	"path/filepath"
	"github.com/eos/tools/errs"
	"gopkg.in/yaml.v2"
)

type Loader struct {
	PathResolver PathResolver 
}

func NewLoader(pathResolver PathResolver) *Loader {
	return &Loader{
		PathResolver: pathResolver,
	}
}

func (c *Loader) InitConfig(config any, configName, configFolderPath string) error {
	configFolderPath, err := c.resolveConfigPath(configName, configFolderPath)
	if err != nil {
		return errs.WrapMsg(err, "resolveConfigPath failed", "configName", configName, "configFolderPath", configFolderPath)
	}

	data, err := os.ReadFile(configFolderPath)
	if err != nil {
		return errs.WrapMsg(err, "ReadFile failed", "configFolderPath", configFolderPath)
	}

	if err = yaml.Unmarshal(data, config); err != nil {
		return errs.WrapMsg(err, "failed to unmarshal config data", "configName", configName)
	}

	return nil
}

func (c *Loader) resolveConfigPath(configName, configFolderPath string) (string, error) {
	if configFolderPath == "" {
		var err error
		configFolderPath, err = c.PathResolver.GetDefaultConfigPath()
		if err != nil {
			return "", errs.WrapMsg(err, "GetDefaultConfigPath failed", "configName", configName)
		}
	}

	configFilePath := filepath.Join(configFolderPath, configName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// Attempt to load from project root if not found in specified path
		projectRoot, err := c.PathResolver.GetProjectRoot()
		if err != nil {
			return "", err
		}
		configFilePath = filepath.Join(projectRoot, "config", configName)
	}
	return configFilePath, nil
}
// Parser Configures the parser interface.
type Parser interface {
	Parse(data []byte, out any) error
}

// YAMLParser Configuration parser in YAML format.
type YAMLParser struct{}

func (y *YAMLParser) Parse(data []byte, out any) error {
	return yaml.Unmarshal(data, out)
}
// ConfigSource configuring source interfaces.
type ConfigSource interface {
	Read() ([]byte, error)
}

// EnvVarSource read a configuration from an environment variable.
type EnvVarSource struct {
	VarName string
}

func (e *EnvVarSource) Read() ([]byte, error) {
	value, exists := os.LookupEnv(e.VarName)
	if !exists {
		return nil, errs.New("environment variable not set").Wrap()
	}
	return []byte(value), nil
}

// FileSystemSource read a configuration from a file.
type FileSystemSource struct {
	FilePath string
}

func (f *FileSystemSource) Read() ([]byte, error) {
	r, err := os.ReadFile(f.FilePath)
	return r, errs.WrapMsg(err, "ReadFile failed ", "FilePath", f.FilePath)
}


type Manager struct {
	sources []ConfigSource
	parser  Parser
}

func NewManager(parser Parser) *Manager {
	return &Manager{
		parser: parser,
	}
}

func (cm *Manager) AddSource(source ConfigSource) {
	cm.sources = append(cm.sources, source)
}

func (cm *Manager) Load(config any) error {
	for _, source := range cm.sources {
		if data, err := source.Read(); err == nil {
			if err := cm.parser.Parse(data, config); err != nil {
				return err
			}
			break
		}
	}
	return nil
}
// PathResolver defines methods for resolving paths related to the application.
type PathResolver interface {
	GetDefaultConfigPath() (string, error)
	GetProjectRoot() (string, error)
}

type defaultPathResolver struct{}

// NewPathResolver creates a new instance of the default path resolver.
func NewPathResolver() *defaultPathResolver {
	return &defaultPathResolver{}
}

func (d *defaultPathResolver) GetDefaultConfigPath(relativePath string) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", errs.WrapMsg(err, "Executable failed")
	}

	configPath := filepath.Join(filepath.Dir(executablePath), relativePath)
	return configPath, nil
}

// GetProjectRoot returns the project's root directory based on the relative depth specified.
// The depth parameter specifies how many levels up from the directory of the executable the project root is located.
func (d *defaultPathResolver) GetProjectRoot(depth int) (string, error) {
	executablePath, err := os.Executable()
	if err != nil {
		return "", errs.WrapMsg(err, "Executable failed")
	}

	// Move up the specified number of directories to find the project root
	projectRoot := executablePath
	for i := 0; i < depth; i++ {
		projectRoot = filepath.Dir(projectRoot)
	}

	return projectRoot, nil
}
