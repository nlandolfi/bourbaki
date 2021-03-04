# generated automatically by sheet.go

make:
	make probability_densities.pdf

probability_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_densities.tex

probability_densities.pdf: ../../*.tex ../../trademark.pdf *.tex probability_densities.tex
	pdflatex probability_densities.tex
	make defs


open: probability_densities.pdf
	open probability_densities.pdf

o:
	make open

defs: probability_densities.tex
	defs probability_densities.tex

spell:
	aspell -c probability_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile