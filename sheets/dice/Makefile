# generated automatically by sheet.go

make:
	make dice.pdf

dice.tex: sheet.tex
	sheet -mode c -in sheet.tex > dice.tex

dice.pdf: ../../*.tex ../../trademark.pdf *.tex dice.tex
	pdflatex dice.tex
	make defs


open: dice.pdf
	open dice.pdf

o:
	make open

defs: dice.tex
	defs dice.tex

spell:
	aspell -c dice.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile