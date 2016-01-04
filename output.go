package dochead

import (
    "io"
    "io/ioutil"
    "text/template"
)

func WriteAPIDefinition(writer io.Writer, apiDefinition ApiDefinition, templateFile string) {  
    tpl, err := ioutil.ReadFile(templateFile)
    processError(err)
    
    parsedTemplate, err := template.New("apiDefinition").Funcs(template.FuncMap{
    "decrement": func(a int) int {
        return a - 1
    },
    }).Parse(string(tpl))
    processError(err)
    
    err = parsedTemplate.Execute(writer, apiDefinition)
    processError(err)
}