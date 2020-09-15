real_length_impossible.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_length_impossible.tex
	make defs

make:
	make real_length_impossible.pdf

open: real_length_impossible.pdf
	open real_length_impossible.pdf

defs: real_length_impossible.tex
	defs real_length_impossible.tex
