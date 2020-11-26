set_notation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex set_notation.tex
	make defs

make:
	make set_notation.pdf

open: set_notation.pdf
	open set_notation.pdf

o:
	make open

defs: set_notation.tex
	defs set_notation.tex

spell:
	aspell -c set_notation.tex

s:
	make spell
