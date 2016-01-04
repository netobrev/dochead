package dochead

import (
	"testing"
)

func assertEquals(t *testing.T, value, expected, name string) {
    if value != expected {
        t.Errorf("api resource %s \"%s\" does not match \"%s\"", name, value, expected)
    }
}

func TestMarkdown(t *testing.T) {
    file := "./parser_test.ogdl"
   
    apiResources, _ := ReadAPIResources(file)
    
    if len(apiResources) != 1 {
        t.Errorf("illegal array length %d", len(apiResources))
    }
    
    apiResource := apiResources[0]
    
    assertEquals(t, apiResource.Name, "Get Balance", "name")
    assertEquals(t, apiResource.Method, "GET", "method")
    assertEquals(t, apiResource.URI, "/balances/{id}", "URI")
}