monotone_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex monotone_algebra.tex
	make defs

make:
	make monotone_algebra.pdf

open: monotone_algebra.pdf
	open monotone_algebra.pdf

defs: monotone_algebra.tex
	defs monotone_algebra.tex
