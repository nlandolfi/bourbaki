product_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex product_measures.tex
	make defs

make:
	make product_measures.pdf

open: product_measures.pdf
	open product_measures.pdf

defs: product_measures.tex
	defs product_measures.tex
