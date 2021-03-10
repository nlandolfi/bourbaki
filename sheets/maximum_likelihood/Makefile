# generated automatically by sheet.go

make:
	make maximum_likelihood.pdf

maximum_likelihood.tex: sheet.tex
	sheet -mode c -in sheet.tex > maximum_likelihood.tex

maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex maximum_likelihood.tex
	pdflatex maximum_likelihood.tex
	make defs


open: maximum_likelihood.pdf
	open maximum_likelihood.pdf

o:
	make open

defs: maximum_likelihood.tex
	defs maximum_likelihood.tex

spell:
	aspell -c maximum_likelihood.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile