vector_bases.tex: sheet.tex
	sheet -mode c -in sheet.tex > vector_bases.tex

vector_bases.pdf: ../../*.tex ../../trademark.pdf *.tex vector_bases.tex
	pdflatex vector_bases.tex
	make defs

make:
	make vector_bases.pdf

open: vector_bases.pdf
	open vector_bases.pdf

o:
	make open

defs: vector_bases.tex
	defs vector_bases.tex

spell:
	aspell -c vector_bases.tex

s:
	make spell