set_examples.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex set_examples.tex
	make defs

make:
	make set_examples.pdf

open: set_examples.pdf
	open set_examples.pdf

o:
	make open

defs: set_examples.tex
	defs set_examples.tex

spell:
	aspell -c set_examples.tex

s:
	make spell
