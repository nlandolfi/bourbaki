arithmetic.pdf: ../*.tex *.tex ../trademark.pdf
	pdflatex arithmetic.tex
	make defs

make:
	make arithmetic.pdf

open: arithmetic.pdf
	open arithmetic.pdf

defs: arithmetic.tex
	defs arithmetic.tex
