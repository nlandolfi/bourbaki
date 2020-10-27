variation_measure.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex variation_measure.tex
	make defs

make:
	make variation_measure.pdf

open: variation_measure.pdf
	open variation_measure.pdf

defs: variation_measure.tex
	defs variation_measure.tex
