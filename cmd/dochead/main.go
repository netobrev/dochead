package main

import (
	"os"
	"fmt"
    "strings"
    "github.com/netobrev/dochead"
    "github.com/jessevdk/go-flags"
)

type options struct {
    InputFiles []string `short:"f" long:"file" description:"Input files" value-name:"FILE"`
    Template string `short:"t" long:"template" description:"Output Template" value-name:"FILE"`
}

func main() {
    o := options {}
	args, err := flags.ParseArgs(&o, os.Args[1:])
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Input Files: %s\nTemplate: %s\n", o.InputFiles, o.Template)
    fmt.Printf("Remaining args: %s\n", strings.Join(args, " "))
    
    if o.Template != "" {
        for _,file := range o.InputFiles {
            fmt.Printf("Processing file: %s\n", file)
            apiDefinition := dochead.ReadAPIDefinition(file)
            dochead.WriteAPIDefinition(apiDefinition, o.Template)
        }
    } else {
        fmt.Printf("No Template specified, nothing to print!\n")
    }
}

