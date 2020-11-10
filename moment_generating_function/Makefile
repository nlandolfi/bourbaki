moment_generating_function.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex moment_generating_function.tex
	make defs

make:
	make moment_generating_function.pdf

open: moment_generating_function.pdf
	open moment_generating_function.pdf

o:
	make open

defs: moment_generating_function.tex
	defs moment_generating_function.tex
