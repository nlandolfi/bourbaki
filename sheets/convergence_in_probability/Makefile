convergence_in_probability.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex convergence_in_probability.tex
	make defs

make:
	make convergence_in_probability.pdf

open: convergence_in_probability.pdf
	open convergence_in_probability.pdf

o:
	make open

defs: convergence_in_probability.tex
	defs convergence_in_probability.tex
