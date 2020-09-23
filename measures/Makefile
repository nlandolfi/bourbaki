measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measures.tex
	make defs

make:
	make measures.pdf

open: measures.pdf
	open measures.pdf

defs: measures.tex
	defs measures.tex
