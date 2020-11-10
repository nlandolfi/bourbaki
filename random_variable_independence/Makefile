random_variable_independence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variable_independence.tex
	make defs

make:
	make random_variable_independence.pdf

open: random_variable_independence.pdf
	open random_variable_independence.pdf

o:
	make open

defs: random_variable_independence.tex
	defs random_variable_independence.tex
