intervals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex intervals.tex
	make defs

make:
	make intervals.pdf

open: intervals.pdf
	open intervals.pdf

defs: intervals.tex
	defs intervals.tex

o:
	make open
