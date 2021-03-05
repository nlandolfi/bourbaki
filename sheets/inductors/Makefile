# generated automatically by sheet.go

make:
	make inductors.pdf

inductors.tex: sheet.tex
	sheet -mode c -in sheet.tex > inductors.tex

inductors.pdf: ../../*.tex ../../trademark.pdf *.tex inductors.tex
	pdflatex inductors.tex
	make defs


open: inductors.pdf
	open inductors.pdf

o:
	make open

defs: inductors.tex
	defs inductors.tex

spell:
	aspell -c inductors.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile