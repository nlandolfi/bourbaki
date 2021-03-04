# generated automatically by sheet.go

make:
	make conditional_densities.pdf

conditional_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > conditional_densities.tex

conditional_densities.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_densities.tex
	pdflatex conditional_densities.tex
	make defs


open: conditional_densities.pdf
	open conditional_densities.pdf

o:
	make open

defs: conditional_densities.tex
	defs conditional_densities.tex

spell:
	aspell -c conditional_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile