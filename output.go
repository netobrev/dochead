package dochead

import (
    "io"
    "io/ioutil"
    "text/template"
)

func WriteAPIDefinition(writer io.Writer, apiDefinition ApiDefinition, templateFile string) {  
    tpl, err := ioutil.ReadFile(templateFile)
    processError(err)
    
    parsedTemplate, err := template.New("apiDefinition").Parse(string(tpl))
    processError(err)
    
    err = parsedTemplate.Execute(writer, apiDefinition)
    processError(err)
}