zero.tex: sheet.tex
	sheet -mode c -in sheet.tex > zero.tex

zero.pdf: ../../*.tex ../../trademark.pdf *.tex zero.tex
	pdflatex zero.tex
	make defs

make:
	make zero.pdf

open: zero.pdf
	open zero.pdf

o:
	make open

defs: zero.tex
	defs zero.tex

spell:
	aspell -c zero.tex

s:
	make spell