cumulative_distribution_function_and_law.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex cumulative_distribution_function_and_law.tex
	make defs

make:
	make cumulative_distribution_function_and_law.pdf

open: cumulative_distribution_function_and_law.pdf
	open cumulative_distribution_function_and_law.pdf

o:
	make open

defs: cumulative_distribution_function_and_law.tex
	defs cumulative_distribution_function_and_law.tex
