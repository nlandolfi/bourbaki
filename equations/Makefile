equations.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex equations.tex
	make defs

make:
	make equations.pdf

open: equations.pdf
	open equations.pdf

defs: equations.tex
	defs equations.tex
