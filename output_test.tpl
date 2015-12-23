# My test resources

This template should be used for the test.
{{ range .Resources }}
## {{ .Name }}
`{{ .URI }}`

{{ end }}