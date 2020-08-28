nets.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex nets.tex
	make defs

make:
	make nets.pdf

open: nets.pdf
	open nets.pdf

defs: nets.tex
	defs nets.tex
