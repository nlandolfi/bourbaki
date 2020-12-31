metrology.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex metrology.tex
	make defs

make:
	make metrology.pdf

open: metrology.pdf
	open metrology.pdf

defs: metrology.tex
	defs metrology.tex
