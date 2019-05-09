(progn
	; define variable in other file
	(require "vars.lisp")
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
	(def q3 (* 3 3))
	(def q4 (* 4 4))
	(print (sqrt (+ q3 q4)))
	; test minus and division
	(print (- 10 6))
	(print (/ 1 2)))
; this is a comment

