natural_induction.pdf: ../format.tex ../macros.tex ../trademark.pdf natural_induction.tex macros.tex
	pdflatex natural_induction.tex
	make defs

make:
	make natural_induction.pdf

open: natural_induction.pdf
	open natural_induction.pdf

o:
	make open

defs: natural_induction.tex
	defs natural_induction.tex
