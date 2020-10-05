partitions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex partitions.tex
	make defs

make:
	make partitions.pdf

open: partitions.pdf
	open partitions.pdf

defs: partitions.tex
	defs partitions.tex
