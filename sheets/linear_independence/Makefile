# generated automatically by sheet.go

make:
	make linear_independence.pdf

linear_independence.tex: sheet.tex
	sheet -mode c -in sheet.tex > linear_independence.tex

linear_independence.pdf: ../../*.tex ../../trademark.pdf *.tex linear_independence.tex
	pdflatex linear_independence.tex
	make defs


open: linear_independence.pdf
	open linear_independence.pdf

o:
	make open

defs: linear_independence.tex
	defs linear_independence.tex

spell:
	aspell -c linear_independence.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile