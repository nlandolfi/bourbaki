cartesian_product.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex cartesian_product.tex
	make defs

make:
	make cartesian_product.pdf

open: cartesian_product.pdf
	open cartesian_product.pdf

defs: cartesian_product.tex
	defs cartesian_product.tex
