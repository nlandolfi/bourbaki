triangle_equality.tex: sheet.tex
	sheet -mode c -in sheet.tex > triangle_equality.tex

triangle_equality.pdf: ../../*.tex ../../trademark.pdf *.tex triangle_equality.tex
	pdflatex triangle_equality.tex
	make defs

make:
	make triangle_equality.pdf

open: triangle_equality.pdf
	open triangle_equality.pdf

o:
	make open

defs: triangle_equality.tex
	defs triangle_equality.tex

spell:
	aspell -c triangle_equality.tex

s:
	make spell