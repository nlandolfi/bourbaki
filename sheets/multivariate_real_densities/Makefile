# generated automatically by sheet.go

make:
	make multivariate_real_densities.pdf

multivariate_real_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > multivariate_real_densities.tex

multivariate_real_densities.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_real_densities.tex
	pdflatex multivariate_real_densities.tex
	make defs


open: multivariate_real_densities.pdf
	open multivariate_real_densities.pdf

o:
	make open

defs: multivariate_real_densities.tex
	defs multivariate_real_densities.tex

spell:
	aspell -c multivariate_real_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile