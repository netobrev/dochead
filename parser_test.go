package dochead

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
    file := "./parser_test.ogdl"
   
    apiDefinition, _ := ReadAPIResources(file)
    fmt.Printf("Api %s\n", apiDefinition)
}