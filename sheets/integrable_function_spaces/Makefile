integrable_function_spaces.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex integrable_function_spaces.tex
	make defs

make:
	make integrable_function_spaces.pdf

open: integrable_function_spaces.pdf
	open integrable_function_spaces.pdf

o:
	make open

defs: integrable_function_spaces.tex
	defs integrable_function_spaces.tex
