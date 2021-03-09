# generated automatically by sheet.go

make:
	make optimization.pdf

optimization.tex: sheet.tex
	sheet -mode c -in sheet.tex > optimization.tex

optimization.pdf: ../../*.tex ../../trademark.pdf *.tex optimization.tex
	pdflatex optimization.tex
	make defs


open: optimization.pdf
	open optimization.pdf

o:
	make open

defs: optimization.tex
	defs optimization.tex

spell:
	aspell -c optimization.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile