# generated automatically by sheet.go

make:
	make probability.pdf

probability.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability.tex

probability.pdf: ../../*.tex ../../trademark.pdf *.tex probability.tex
	pdflatex probability.tex
	make defs


open: probability.pdf
	open probability.pdf

o:
	make open

defs: probability.tex
	defs probability.tex

spell:
	aspell -c probability.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile