(progn
	; define variable in other file
	(import "vars.lisp")
	(print chinese)
	; print some string
  (print "hello")
  (print "world")
	; print chinese string
	(print "你好，世界！")
	; test plus
  (print (+ 10 20))
	(print "sqrt(3*3 + 4*4) = ")
	; test multiply and plus
	(defvar q3 (* 3 3))
	(defvar q4 (* 4 4))
	(print (sqrt (+ q3 q4)))
	; test minus and division
	(print (- 10 6))
	(print (/ 1 2))
	; test lambda
	(defvar say-hello
		(lambda (name) 
      (progn 
        (print "hello from lambda")
        (print name))))
	; call lambda
	(say-hello "nicolas"))
; this is a comment

