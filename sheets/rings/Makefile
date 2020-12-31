rings.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex rings.tex
	make defs

make:
	make rings.pdf

open: rings.pdf
	open rings.pdf

defs: rings.tex
	defs rings.tex
