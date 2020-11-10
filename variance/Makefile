variance.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex variance.tex
	make defs

make:
	make variance.pdf

open: variance.pdf
	open variance.pdf

o:
	make open

defs: variance.tex
	defs variance.tex
