product_sigma_algebras.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex product_sigma_algebras.tex
	make defs

make:
	make product_sigma_algebras.pdf

open: product_sigma_algebras.pdf
	open product_sigma_algebras.pdf

defs: product_sigma_algebras.tex
	defs product_sigma_algebras.tex
