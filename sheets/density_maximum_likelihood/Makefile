# generated automatically by sheet.go

make:
	make density_maximum_likelihood.pdf

density_maximum_likelihood.tex: sheet.tex
	sheet -mode c -in sheet.tex > density_maximum_likelihood.tex

density_maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex density_maximum_likelihood.tex
	pdflatex density_maximum_likelihood.tex
	make defs


open: density_maximum_likelihood.pdf
	open density_maximum_likelihood.pdf

o:
	make open

defs: density_maximum_likelihood.tex
	defs density_maximum_likelihood.tex

spell:
	aspell -c density_maximum_likelihood.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile