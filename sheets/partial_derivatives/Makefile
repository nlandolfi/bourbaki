# generated automatically by sheet.go

make:
	make partial_derivatives.pdf

partial_derivatives.tex: sheet.tex
	sheet -mode c -in sheet.tex > partial_derivatives.tex

partial_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex partial_derivatives.tex
	pdflatex partial_derivatives.tex
	make defs


open: partial_derivatives.pdf
	open partial_derivatives.pdf

o:
	make open

defs: partial_derivatives.tex
	defs partial_derivatives.tex

spell:
	aspell -c partial_derivatives.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile