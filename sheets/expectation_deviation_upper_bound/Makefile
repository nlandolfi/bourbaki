expectation_deviation_upper_bound.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex expectation_deviation_upper_bound.tex
	make defs

make:
	make expectation_deviation_upper_bound.pdf

open: expectation_deviation_upper_bound.pdf
	open expectation_deviation_upper_bound.pdf

o:
	make open

defs: expectation_deviation_upper_bound.tex
	defs expectation_deviation_upper_bound.tex
