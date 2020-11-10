product_under_independence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex product_under_independence.tex
	make defs

make:
	make product_under_independence.pdf

open: product_under_independence.pdf
	open product_under_independence.pdf

o:
	make open

defs: product_under_independence.tex
	defs product_under_independence.tex
