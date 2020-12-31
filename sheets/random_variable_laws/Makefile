random_variable_laws.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variable_laws.tex
	make defs

make:
	make random_variable_laws.pdf

open: random_variable_laws.pdf
	open random_variable_laws.pdf

defs: random_variable_laws.tex
	defs random_variable_laws.tex
