# generated automatically by sheet.go

make:
	make coins.pdf

coins.tex: sheet.tex
	sheet -mode c -in sheet.tex > coins.tex

coins.pdf: ../../*.tex ../../trademark.pdf *.tex coins.tex
	pdflatex coins.tex
	make defs


open: coins.pdf
	open coins.pdf

o:
	make open

defs: coins.tex
	defs coins.tex

spell:
	aspell -c coins.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile