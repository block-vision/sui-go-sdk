package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type KeyStoreConfig struct {
	Keys []string
}

var keyStoreConfig KeyStoreConfig

func GetKeyStore(path string) (KeyStoreConfig, error) {
	err := loadConfig(path)
	return keyStoreConfig, err
}

func loadConfig(dir string) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &keyStoreConfig.Keys)
	if err != nil {
		return err
	}
	return nil
}
