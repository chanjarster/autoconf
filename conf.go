package autoconf

import (
	"flag"
	"fmt"
	"os"
)

func Load(p interface{}, yamlFlagName string) {

	initStruct(p)

	resolvers := make([]ConfResolver, 0)

	yamlFile := ""
	if yamlFlagName != "" {
		flag.StringVar(&yamlFile, yamlFlagName, "", "yaml config file")
	}

	er := &envResolver{}
	er.init(p)

	fr := &flagResolver{}
	fr.flagSet = flag.CommandLine
	fr.init(p)

	flag.Parse()

	if yamlFile != "" {
		yr := &yamlFileResolver{File: yamlFile}
		resolvers = append(resolvers, yr)
	}
	resolvers = append(resolvers, er)
	resolvers = append(resolvers, fr)

	for _, resolver := range resolvers {
		if err := resolver.Resolve(p); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
	}

}
