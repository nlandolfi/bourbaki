# generated automatically by sheet.go

make:
	make gaussian_maximum_likelihood.pdf

gaussian_maximum_likelihood.tex: sheet.tex
	sheet -mode c -in sheet.tex > gaussian_maximum_likelihood.tex

gaussian_maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex gaussian_maximum_likelihood.tex
	pdflatex gaussian_maximum_likelihood.tex
	make defs


open: gaussian_maximum_likelihood.pdf
	open gaussian_maximum_likelihood.pdf

o:
	make open

defs: gaussian_maximum_likelihood.tex
	defs gaussian_maximum_likelihood.tex

spell:
	aspell -c gaussian_maximum_likelihood.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile