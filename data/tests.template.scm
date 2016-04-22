(load "testing.scm")


{{range .Problems}}
(define (test-{{.}}))
{{end}}