set_builder_notation.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex set_builder_notation.tex
	make defs

make:
	make set_builder_notation.pdf

open: set_builder_notation.pdf
	open set_builder_notation.pdf

o:
	make open

defs: set_builder_notation.tex
	defs set_builder_notation.tex

spell:
	aspell -c set_builder_notation.tex

s:
	make spell
