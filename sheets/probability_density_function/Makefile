probability_density_function.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex probability_density_function.tex
	make defs

make:
	make probability_density_function.pdf

open: probability_density_function.pdf
	open probability_density_function.pdf

o:
	make open

defs: probability_density_function.tex
	defs probability_density_function.tex
