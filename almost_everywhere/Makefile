almost_everywhere.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex almost_everywhere.tex
	make defs

make:
	make almost_everywhere.pdf

open: almost_everywhere.pdf
	open almost_everywhere.pdf

defs: almost_everywhere.tex
	defs almost_everywhere.tex
