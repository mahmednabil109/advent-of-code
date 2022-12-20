(load "~/quicklisp/setup.lisp")
(ql:quickload "str")

(format t 
	"~d~%"
	(loop for item in 
		(mapcar
			(lambda 
				(elfs)
				(apply
					'+
					(mapcar 
						(lambda (cal) (parse-integer cal))
						(str:split " " (str:remove-punctuation elfs))
					)
				)
			)
			(str:split
				(str:repeat 2 (string #\newline))
				(str:from-file "./input1")
			)
		)
		maximizing item
	)
)

