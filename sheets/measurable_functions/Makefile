measurable_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measurable_functions.tex
	make defs

make:
	make measurable_functions.pdf

open: measurable_functions.pdf
	open measurable_functions.pdf

o:
	make open

defs: measurable_functions.tex
	defs measurable_functions.tex
