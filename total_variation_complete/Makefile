total_variation_complete.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex total_variation_complete.tex
	make defs

make:
	make total_variation_complete.pdf

open: total_variation_complete.pdf
	open total_variation_complete.pdf

defs: total_variation_complete.tex
	defs total_variation_complete.tex
