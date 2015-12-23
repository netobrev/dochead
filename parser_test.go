package dochead

import (
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
    file := "./parser_test.md"
   
    apiDefinition := ReadAPIDefinition(file)
    fmt.Printf("Api %s\n", apiDefinition)
}