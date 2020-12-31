absolute_value.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex absolute_value.tex
	make defs

make:
	make absolute_value.pdf

open: absolute_value.pdf
	open absolute_value.pdf

defs: absolute_value.tex
	defs absolute_value.tex
