package config

type Manager struct {
	sources []ConfigSource
	parser Parser
}
func NewManager(parser Parser) *Manager {
	return &Manager{
		parser : parser,
	}
}	

