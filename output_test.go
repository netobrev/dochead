package dochead

import (
	"testing"
)

func TestOutput(t *testing.T) {
    firstAPIResouse := ApiResource {"First", "GET", "/first", "<h1>First</h1>"}
    secondAPIResource := ApiResource {"Second", "GET", "/second", "<h1>Second</h1>"} 
   
    apiDefinition := ApiDefinition { []ApiResource { firstAPIResouse, secondAPIResource } }
   
    WriteAPIDefinition(apiDefinition, "./output_test.tpl")
}