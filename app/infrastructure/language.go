package infrastructure

import (
	"io/ioutil"

	"golang.org/x/text/message"
	"gopkg.in/yaml.v2"
)

func NewInitLanguage(c *Config) {
	for i := 0; i < len(c.Language); i++ {
		f, err := ioutil.ReadFile("./languages/" + c.Language[i].String() + ".yaml")
		if err != nil {
			panic(err.Error())
		}

		var m map[string]string
		err = yaml.Unmarshal([]byte(f), &m)
		if err != nil {
			panic(err.Error())
		}

		for k, v := range m {
			err = message.SetString(c.Language[i], k, v)
			if err != nil {
				panic(err.Error())
			}
		}
	}
}
