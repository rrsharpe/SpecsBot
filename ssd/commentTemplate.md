{{ if not .Confident }}*I am not confident about this result. Is the brand correct?*{{end}}

### {{ .Title }}

{{ if .Interface }} * Interface: **{{ .Interface }}** {{end}}
{{ if .FormFactor }} * Form Factor: **{{ .FormFactor }}** {{end}}
{{ if .Capacities }} * Capacities: **{{ .Capacities }}** {{end}}
{{ if .Controller }} * Controller: **{{ .Controller }}** {{end}}
{{ if .Configuration }} * Configuration: **{{ .Configuration }}** {{end}}
{{ if .DRAM }} * DRAM: **{{ .DRAM }}** {{end}}
{{ if .HMB }} * HMB: **{{ .HMB }}** {{end}}
{{ if .NANDBrand }} * NAND Brand: **{{ .NANDBrand }}** {{end}}
{{ if .NANDType }} * NAND Type: **{{ .NANDType }}** {{end}}
{{ if .Layers }} * Layers: **{{ .Layers }}** {{end}}
{{ if .ReadWrite }} * Read/Write: **{{ .ReadWrite }}** {{end}}
{{ if .Categories }} * Categories: **{{ .Categories }}** {{end}}
{{ if .Notes }} * Notes: **{{ .Notes }}** {{end}}
{{ if .AltNames }} * Other Names: {{range $i, $el := .AltNames}}**{{if $i}}, {{end}}{{ $el }}**{{end}}{{end}}

---
^(Inspired by a similar bot in /r/buildapcsales/.)
^(Info is sourced from )[^(NewMaxx's spreadsheet.)](https://docs.google.com/spreadsheets/d/1B27_j9NDPU3cNlj2HKcrfpJKHkOf-Oi1DbuuQva2gT4/edit#gid=0)

^(If I fetched the wrong result please )[^(DM me)](https://www.reddit.com/message/compose/?to=SpecsBot)^( so I can improve my pattern matching.)
