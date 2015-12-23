package dochead

import (
	"testing"
)

func TestOutput(t *testing.T) {
    firstAPIResouse := ApiResource {"First", "GET first"}
    secondAPIResource := ApiResource {"Second", "GET second"} 
   
    apiDefinition := ApiDefinition { []ApiResource { firstAPIResouse, secondAPIResource } }
   
    WriteAPIDefinition(apiDefinition, "./output_test.tpl")
}