package dochead

import (
	"os"
    "io/ioutil"
    "text/template"
)

func WriteAPIDefinition(apiDefinition ApiDefinition, templateFile string) {  
    tpl, err := ioutil.ReadFile(templateFile)
    processError(err)
    
    parsedTemplate, err := template.New("apiDefinition").Parse(string(tpl))
    processError(err)
    
    err = parsedTemplate.Execute(os.Stdout, apiDefinition)
    processError(err)
}