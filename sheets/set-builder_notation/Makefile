set-builder_notation.tex: sheet.tex
	sheet -mode c -in sheet.tex > set-builder_notation.tex

set-builder_notation.pdf: ../../*.tex ../../trademark.pdf *.tex set-builder_notation.tex
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