element_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex element_functions.tex
	make defs

make:
	make element_functions.pdf

open: element_functions.pdf
	open element_functions.pdf

o:
	make open

defs: element_functions.tex
	defs element_functions.tex

spell:
	aspell -c element_functions.tex

s:
	make spell
