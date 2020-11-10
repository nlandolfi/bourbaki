random_variables.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variables.tex
	make defs

make:
	make random_variables.pdf

open: random_variables.pdf
	open random_variables.pdf

o:
	make open

defs: random_variables.tex
	defs random_variables.tex
