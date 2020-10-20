natural_order.pdf: ../format.tex ../macros.tex ../trademark.pdf natural_order.tex macros.tex
	pdflatex natural_order.tex
	make defs

make:
	make natural_order.pdf

open: natural_order.pdf
	open natural_order.pdf

defs: natural_order.tex
	defs natural_order.tex
