measure_derivatives.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measure_derivatives.tex
	make defs

make:
	make measure_derivatives.pdf

open: measure_derivatives.pdf
	open measure_derivatives.pdf

o:
	make open

defs: measure_derivatives.tex
	defs measure_derivatives.tex
