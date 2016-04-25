(load "testing.scm")



{{range .Problems}}
(define (test-exercise-{{.Chapter}}-{{.Number}}))
{{end}}
