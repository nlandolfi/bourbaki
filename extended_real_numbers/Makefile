extended_real_numbers.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex extended_real_numbers.tex
	make defs

make:
	make extended_real_numbers.pdf

open: extended_real_numbers.pdf
	open extended_real_numbers.pdf

defs: extended_real_numbers.tex
	defs extended_real_numbers.tex
