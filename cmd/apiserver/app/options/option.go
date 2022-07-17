package options

import (
	"encoding/json"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/huangjiasingle/suyi/pkg/config/generic"
)

type ServerRunOptions struct {
	Mysql  *generic.MysqlOptions  `json:"mysql,omitempty" yaml:"mysql,omitempty"`
	Redis  *generic.RedisOptions  `json:"redis,omitempty" yaml:"redis,omitempty"`
	System *generic.SystemOptions `json:"system,omitempty" yaml:"system,omitempty"`
	Config string
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Mysql: &generic.MysqlOptions{},
	}
}

func (s *ServerRunOptions) LoadConfig() error {
	content, err := os.ReadFile(s.Config)
	if err != nil {
		return err
	}
	if strings.HasSuffix(s.Config, ".json") {
		if err := json.Unmarshal(content, s); err != nil {
			return err
		}
	}
	if strings.HasSuffix(s.Config, ".yaml") || strings.HasSuffix(s.Config, ".yml") {
		if err := yaml.Unmarshal(content, s); err != nil {
			return err
		}
	}
	return nil
}
