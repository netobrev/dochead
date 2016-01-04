package dochead

import (
    "bytes"
	"testing"
)

const expectedString =
`# My test resources

This template should be used for the test.

## GET First
` + "`/first`" 

func TestOutput(t *testing.T) {
    apiResource := ApiResource {
        "First", 
        "GET", 
        "/first", 
        "Foo", 
        []Parameter {}, 
        Body {}, 
        Return {}, 
        []Example {},
    }

    apiDefinition := ApiDefinition { []ApiResource { apiResource } }
  
    var buffer bytes.Buffer
    WriteAPIDefinition(&buffer, apiDefinition, "./output_test.tpl")
    
    output := buffer.String()
    
    if output != expectedString {
        t.Errorf("output //\n\n%s\n\n// does not match expected template output //\n\n%s\n\n//", output, expectedString)
    }
}