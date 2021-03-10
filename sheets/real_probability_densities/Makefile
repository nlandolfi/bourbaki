# generated automatically by sheet.go

make:
	make real_probability_densities.pdf

real_probability_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_probability_densities.tex

real_probability_densities.pdf: ../../*.tex ../../trademark.pdf *.tex real_probability_densities.tex
	pdflatex real_probability_densities.tex
	make defs


open: real_probability_densities.pdf
	open real_probability_densities.pdf

o:
	make open

defs: real_probability_densities.tex
	defs real_probability_densities.tex

spell:
	aspell -c real_probability_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile