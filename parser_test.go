package dochead

import (
	"fmt"
	"testing"
)

func TestParseURI(t *testing.T) {
    httpURI := "GET /test"
    
    verb, uri := parseURI(httpURI)
    fmt.Printf("Verb: %s, URI: %s", verb, uri)
    if verb != "GET" || uri != "/test" {
        t.Fatalf("could not parse URI %s\n", httpURI)
    }
}

func TestMarkdown(t *testing.T) {
    file := "./parser_test.md"
   
    apiDefinition := ReadAPIDefinition(file)
    fmt.Printf("Api %s\n", apiDefinition)
}