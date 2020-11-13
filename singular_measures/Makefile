singular_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex singular_measures.tex
	make defs

make:
	make singular_measures.pdf

open: singular_measures.pdf
	open singular_measures.pdf

o:
	make open

defs: singular_measures.tex
	defs singular_measures.tex
