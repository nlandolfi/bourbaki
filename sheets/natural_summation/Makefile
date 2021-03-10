# generated automatically by sheet.go

make:
	make natural_summation.pdf

natural_summation.tex: sheet.tex
	sheet -mode c -in sheet.tex > natural_summation.tex

natural_summation.pdf: ../../*.tex ../../trademark.pdf *.tex natural_summation.tex
	pdflatex natural_summation.tex
	make defs


open: natural_summation.pdf
	open natural_summation.pdf

o:
	make open

defs: natural_summation.tex
	defs natural_summation.tex

spell:
	aspell -c natural_summation.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile