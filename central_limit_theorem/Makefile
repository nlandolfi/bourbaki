central_limit_theorem.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex central_limit_theorem.tex
	make defs

make:
	make central_limit_theorem.pdf

open: central_limit_theorem.pdf
	open central_limit_theorem.pdf

o:
	make open

defs: central_limit_theorem.tex
	defs central_limit_theorem.tex
