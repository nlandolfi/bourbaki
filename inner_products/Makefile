inner_products.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex inner_products.tex
	make defs

make:
	make inner_products.pdf

open: inner_products.pdf
	open inner_products.pdf

defs: inner_products.tex
	defs inner_products.tex
