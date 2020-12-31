expectation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex expectation.tex
	make defs

make:
	make expectation.pdf

open: expectation.pdf
	open expectation.pdf

o:
	make open

defs: expectation.tex
	defs expectation.tex
