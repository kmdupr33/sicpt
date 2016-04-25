(load "{{.TestsFile}}")
(load "{{.DepsFile}}")

{{range .Problems}}
;;; Exercise {{.Chapter}}.{{.Number}}

(test-exercise-{{.Chapter}}-{{.Number}})

{{end}}