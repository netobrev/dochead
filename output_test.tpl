# My test resources

This template should be used for the test.
{{ range .Resources }}
## {{ .Verb }} {{ .Name }}
`{{ .URI }}`

{{ end }}