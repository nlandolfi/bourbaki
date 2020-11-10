cumulative_distribution_function_from_density.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex cumulative_distribution_function_from_density.tex
	make defs

make:
	make cumulative_distribution_function_from_density.pdf

open: cumulative_distribution_function_from_density.pdf
	open cumulative_distribution_function_from_density.pdf

o:
	make open

defs: cumulative_distribution_function_from_density.tex
	defs cumulative_distribution_function_from_density.tex
