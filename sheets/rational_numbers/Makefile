rational_numbers.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex rational_numbers.tex
	make defs

make:
	make rational_numbers.pdf

open: rational_numbers.pdf
	open rational_numbers.pdf

defs: rational_numbers.tex
	defs rational_numbers.tex
