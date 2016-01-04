{
  "resources": [
  
    {{ $resources := .Resources }}
    {{ range $resourceIndex, $resource := .Resources }}
    {
        {{ if $resource.Name }}"title": "{{ $resource.Name }}"{{ end }}
        {{ if $resource.Method }}"verb": "{{ $resource.Method }}"{{ end }}
        {{ if $resource.URI }}"uri": "{{ $resource.URI }}"{{ end }}
        {{ if $resource.Description }}"raw": "{{ $resource.Description }}"{{ end }}
        
        {{ if $resource.Parameters }}
        "parameters" : {
            "path" : [
            {{ range $parameterIndex, $parameter := $resource.Parameters }}
                {
                    {{ if $parameter.Name }} "title": "{{ $parameter.Name }}" {{ end }}
                    {{ if $parameter.Type }} "type": "{{ $parameter.Type }}" {{ end }}
                    {{ if $parameter.Description }} "title": "{{ $parameter.Description }}" {{ end }}
                } {{ if lt $parameterIndex (decrement (len $resource.Parameters)) }},{{ end }}
            {{ end }}
            ]
        },
        {{ end }}
        
        {{ if $resource.Body }}
        "body" : {
            {{ if $resource.Body.Accept }}"accept": "{{ $resource.Body.Accept }}"{{ end }}
            {{ if $resource.Body.Schema }}"schema": "{{ $resource.Body.Schema }}"{{ end }}
        },
        {{ end }}
        
        {{ if $resource.Return }}
        "returns" : {
            {{ if $resource.Return.ContentType }}"content_type": "{{ $resource.Return.ContentType }}"{{ end }}
            {{ if $resource.Return.Schema }}"schema": "{{ $resource.Return.Schema }}"{{ end }}
        },
        {{ end }}
        
        {{ if $resource.Examples }}
        "examples" : [
            {{ range $exampleIndex, $example := $resource.Examples }}
            {
                {{ if $example.Name }}"name": "{{ $example.Name }}"{{ end }}
                {{ if $example.Request }}"request": "{{ $example.Request }}"{{ end }}
                {{ if $example.Response }}"response": "{{ $example.Response }}"{{ end }}
            } {{ if lt $exampleIndex (decrement (len $resource.Examples)) }},{{ end }}
            {{ end }}
        ]
        {{ end }}
    } {{ if lt $resourceIndex (decrement (len $resources)) }},{{ end }}
    {{ end }}
}
