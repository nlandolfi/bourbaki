natural_numbers.pdf: ../format.tex ../macros.tex ../trademark.pdf natural_numbers.tex macros.tex
	pdflatex natural_numbers.tex
	make defs

make:
	make natural_numbers.pdf

open: natural_numbers.pdf
	open natural_numbers.pdf

defs: natural_numbers.tex
	defs natural_numbers.tex
