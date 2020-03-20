package autoconf

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type yamlFileResolver struct {
	File string
}

func (y *yamlFileResolver) Resolve(p interface{}) error {
	bytes, err := ioutil.ReadFile(y.File)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, p)
	if err != nil {
		return err
	}
	return nil

}
