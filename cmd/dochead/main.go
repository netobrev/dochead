package main

import (
	"os"
    "github.com/netobrev/dochead"
    "github.com/jessevdk/go-flags"
)

// command line options
type options struct {
    InputFiles []string `short:"f" long:"file" description:"Input files" value-name:"<FILE>" required:"true"`
    Template string `short:"t" long:"template" description:"Output Template" value-name:"<FILE>" required:"true"`
}

func main() {
    o := options {}
	_, err := flags.ParseArgs(&o, os.Args[1:])
    if err != nil {
        os.Exit(1)
    }

    // TODO aggregate input files to a common api definition
    for _, file := range o.InputFiles {
        apiResources, _ := dochead.ReadAPIResources(file)
        apiDefinition := dochead.ApiDefinition{apiResources}
        dochead.WriteAPIDefinition(apiDefinition, o.Template)
    }
}

