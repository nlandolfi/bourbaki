probability_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex probability_measures.tex
	make defs

make:
	make probability_measures.pdf

open: probability_measures.pdf
	open probability_measures.pdf

defs: probability_measures.tex
	defs probability_measures.tex
