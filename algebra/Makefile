algebra.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex algebra.tex
	make defs

make:
	make algebra.pdf

open: algebra.pdf
	open algebra.pdf

defs: algebra.tex
	defs algebra.tex
