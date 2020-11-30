probability.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex probability.tex
	make defs

make:
	make probability.pdf

open: probability.pdf
	open probability.pdf

o:
	make open

defs: probability.tex
	defs probability.tex

spell:
	aspell -c probability.tex

s:
	make spell
