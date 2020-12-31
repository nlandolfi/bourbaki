subsets.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex subsets.tex
	make defs

make:
	make subsets.pdf

open: subsets.pdf
	open subsets.pdf

o:
	make open

defs: subsets.tex
	defs subsets.tex
