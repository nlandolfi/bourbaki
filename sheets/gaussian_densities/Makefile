# generated automatically by sheet.go

make:
	make gaussian_densities.pdf

gaussian_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > gaussian_densities.tex

gaussian_densities.pdf: ../../*.tex ../../trademark.pdf *.tex gaussian_densities.tex
	pdflatex gaussian_densities.tex
	make defs


open: gaussian_densities.pdf
	open gaussian_densities.pdf

o:
	make open

defs: gaussian_densities.tex
	defs gaussian_densities.tex

spell:
	aspell -c gaussian_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile