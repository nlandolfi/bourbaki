set-builder_notation.pdf: ../../*.tex ../../trademark.pdf *.tex
	pdflatex set-builder_notation.tex
	make defs

make:
	make set-builder_notation.pdf

open: set-builder_notation.pdf
	open set-builder_notation.pdf

o:
	make open

defs: set-builder_notation.tex
	defs set-builder_notation.tex

spell:
	aspell -c set-builder_notation.tex

s:
	make spell
