convergence_in_measure.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex convergence_in_measure.tex
	make defs

make:
	make convergence_in_measure.pdf

open: convergence_in_measure.pdf
	open convergence_in_measure.pdf

o:
	make open

defs: convergence_in_measure.tex
	defs convergence_in_measure.tex
