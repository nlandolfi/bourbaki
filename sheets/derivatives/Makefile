# generated automatically by sheet.go

make:
	make derivatives.pdf

derivatives.tex: sheet.tex
	sheet -mode c -in sheet.tex > derivatives.tex

derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex derivatives.tex
	pdflatex derivatives.tex
	make defs


open: derivatives.pdf
	open derivatives.pdf

o:
	make open

defs: derivatives.tex
	defs derivatives.tex

spell:
	aspell -c derivatives.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile