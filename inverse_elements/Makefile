inverse_elements.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex inverse_elements.tex
	make defs

make:
	make inverse_elements.pdf

open: inverse_elements.pdf
	open inverse_elements.pdf

o:
	make open

defs: inverse_elements.tex
	defs inverse_elements.tex

spell:
	aspell -c inverse_elements.tex

s:
	make spell
