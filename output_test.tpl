# My test resources

This template should be used for the test.
{{ range .Resources }}
## {{ .Method }} {{ .Name }}
`{{ .URI }}`{{ end }}