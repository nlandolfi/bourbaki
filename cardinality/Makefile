cardinality.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex cardinality.tex
	make defs

make:
	make cardinality.pdf

open: cardinality.pdf
	open cardinality.pdf

defs: cardinality.tex
	defs cardinality.tex
