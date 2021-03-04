# generated automatically by sheet.go

make:
	make multivariate_gaussians.pdf

multivariate_gaussians.tex: sheet.tex
	sheet -mode c -in sheet.tex > multivariate_gaussians.tex

multivariate_gaussians.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_gaussians.tex
	pdflatex multivariate_gaussians.tex
	make defs


open: multivariate_gaussians.pdf
	open multivariate_gaussians.pdf

o:
	make open

defs: multivariate_gaussians.tex
	defs multivariate_gaussians.tex

spell:
	aspell -c multivariate_gaussians.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile