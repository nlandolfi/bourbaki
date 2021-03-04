# generated automatically by sheet.go

make:
	make fair_coins.pdf

fair_coins.tex: sheet.tex
	sheet -mode c -in sheet.tex > fair_coins.tex

fair_coins.pdf: ../../*.tex ../../trademark.pdf *.tex fair_coins.tex
	pdflatex fair_coins.tex
	make defs


open: fair_coins.pdf
	open fair_coins.pdf

o:
	make open

defs: fair_coins.tex
	defs fair_coins.tex

spell:
	aspell -c fair_coins.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile