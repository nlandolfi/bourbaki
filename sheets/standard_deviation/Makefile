standard_deviation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex standard_deviation.tex
	make defs

make:
	make standard_deviation.pdf

open: standard_deviation.pdf
	open standard_deviation.pdf

o:
	make open

defs: standard_deviation.tex
	defs standard_deviation.tex
