random_variable_sigma_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex random_variable_sigma_algebra.tex
	make defs

make:
	make random_variable_sigma_algebra.pdf

open: random_variable_sigma_algebra.pdf
	open random_variable_sigma_algebra.pdf

o:
	make open

defs: random_variable_sigma_algebra.tex
	defs random_variable_sigma_algebra.tex
