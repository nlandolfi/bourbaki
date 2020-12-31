simple_integral_monotonicity.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex simple_integral_monotonicity.tex
	make defs

make:
	make simple_integral_monotonicity.pdf

open: simple_integral_monotonicity.pdf
	open simple_integral_monotonicity.pdf

defs: simple_integral_monotonicity.tex
	defs simple_integral_monotonicity.tex
