direct_products.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex direct_products.tex
	make defs

make:
	make direct_products.pdf

open: direct_products.pdf
	open direct_products.pdf

defs: direct_products.tex
	defs direct_products.tex
