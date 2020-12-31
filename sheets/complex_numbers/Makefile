complex_numbers.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex complex_numbers.tex
	make defs

make:
	make complex_numbers.pdf

open: complex_numbers.pdf
	open complex_numbers.pdf

o:
	make open

defs: complex_numbers.tex
	defs complex_numbers.tex
