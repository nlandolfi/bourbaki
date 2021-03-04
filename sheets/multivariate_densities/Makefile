# generated automatically by sheet.go

make:
	make multivariate_densities.pdf

multivariate_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > multivariate_densities.tex

multivariate_densities.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_densities.tex
	pdflatex multivariate_densities.tex
	make defs


open: multivariate_densities.pdf
	open multivariate_densities.pdf

o:
	make open

defs: multivariate_densities.tex
	defs multivariate_densities.tex

spell:
	aspell -c multivariate_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile