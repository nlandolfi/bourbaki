random_variable_moments.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variable_moments.tex
	make defs

make:
	make random_variable_moments.pdf

open: random_variable_moments.pdf
	open random_variable_moments.pdf

o:
	make open

defs: random_variable_moments.tex
	defs random_variable_moments.tex
