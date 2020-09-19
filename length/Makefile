length.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex length.tex
	make defs

make:
	make length.pdf

open: length.pdf
	open length.pdf

defs: length.tex
	defs length.tex
