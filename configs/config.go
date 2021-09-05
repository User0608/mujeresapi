package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cors struct {
	AllowOrigins []string `json:"allow_origins"`
	AllowMethods []string `json:"allow_methods"`
}
type Certificates struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

type Config struct {
	Address            string       `json:"address"`
	MediaRootDir       string       `json:"media_dir"`
	PublicMediaDirPath string       `json:"public_media_dir_path"`
	Cors               Cors         `json:"CORS"`
	Certificates       Certificates `json:"certificates"`
}

func LoadConfigs(path string) (*Config, error) {
	configs := &Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("No se pudo abrir el archivo de configuraciones: %s", err.Error())
	}
	if err := json.Unmarshal(file, configs); err != nil {
		return nil, fmt.Errorf("Archivo de configuraciones incorrecto, ERR: %s", err.Error())
	}
	return configs, nil
}
