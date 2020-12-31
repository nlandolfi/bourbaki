monotone_classes.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex monotone_classes.tex
	make defs

make:
	make monotone_classes.pdf

open: monotone_classes.pdf
	open monotone_classes.pdf

defs: monotone_classes.tex
	defs monotone_classes.tex
