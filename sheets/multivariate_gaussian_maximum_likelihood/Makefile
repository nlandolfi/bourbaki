# generated automatically by sheet.go

make:
	make multivariate_gaussian_maximum_likelihood.pdf

multivariate_gaussian_maximum_likelihood.tex: sheet.tex
	sheet -mode c -in sheet.tex > multivariate_gaussian_maximum_likelihood.tex

multivariate_gaussian_maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_gaussian_maximum_likelihood.tex
	pdflatex multivariate_gaussian_maximum_likelihood.tex
	make defs


open: multivariate_gaussian_maximum_likelihood.pdf
	open multivariate_gaussian_maximum_likelihood.pdf

o:
	make open

defs: multivariate_gaussian_maximum_likelihood.tex
	defs multivariate_gaussian_maximum_likelihood.tex

spell:
	aspell -c multivariate_gaussian_maximum_likelihood.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile