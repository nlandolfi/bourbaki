real_numbers.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_numbers.tex
	make defs

make:
	make real_numbers.pdf

open: real_numbers.pdf
	open real_numbers.pdf

defs: real_numbers.tex
	defs real_numbers.tex
