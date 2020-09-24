finite_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex finite_measures.tex
	make defs

make:
	make finite_measures.pdf

open: finite_measures.pdf
	open finite_measures.pdf

defs: finite_measures.tex
	defs finite_measures.tex
