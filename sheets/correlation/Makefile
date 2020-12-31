correlation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex correlation.tex
	make defs

make:
	make correlation.pdf

open: correlation.pdf
	open correlation.pdf

o:
	make open

defs: correlation.tex
	defs correlation.tex
