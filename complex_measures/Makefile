complex_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex complex_measures.tex
	make defs

make:
	make complex_measures.pdf

open: complex_measures.pdf
	open complex_measures.pdf

o:
	make open

defs: complex_measures.tex
	defs complex_measures.tex
