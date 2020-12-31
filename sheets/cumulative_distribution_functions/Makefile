cumulative_distribution_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex cumulative_distribution_functions.tex
	make defs

make:
	make cumulative_distribution_functions.pdf

open: cumulative_distribution_functions.pdf
	open cumulative_distribution_functions.pdf

o:
	make open

defs: cumulative_distribution_functions.tex
	defs cumulative_distribution_functions.tex
