zero.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex zero.tex
	make defs

make:
	make zero.pdf

open: zero.pdf
	open zero.pdf

defs: zero.tex
	defs zero.tex
