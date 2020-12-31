total_variation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex total_variation.tex
	make defs

make:
	make total_variation.pdf

open: total_variation.pdf
	open total_variation.pdf

defs: total_variation.tex
	defs total_variation.tex
