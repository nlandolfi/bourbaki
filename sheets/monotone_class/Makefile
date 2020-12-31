monotone_class.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex monotone_class.tex
	make defs

make:
	make monotone_class.pdf

open: monotone_class.pdf
	open monotone_class.pdf

defs: monotone_class.tex
	defs monotone_class.tex
