probability_space.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex probability_space.tex
	make defs

make:
	make probability_space.pdf

open: probability_space.pdf
	open probability_space.pdf

defs: probability_space.tex
	defs probability_space.tex
