simple_integral_additivity.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex simple_integral_additivity.tex
	make defs

make:
	make simple_integral_additivity.pdf

open: simple_integral_additivity.pdf
	open simple_integral_additivity.pdf

defs: simple_integral_additivity.tex
	defs simple_integral_additivity.tex
