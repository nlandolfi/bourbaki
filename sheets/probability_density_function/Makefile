# generated automatically by sheet.go

make:
	make probability_density_function.pdf

probability_density_function.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_density_function.tex

probability_density_function.pdf: ../../*.tex ../../trademark.pdf *.tex probability_density_function.tex
	pdflatex probability_density_function.tex
	make defs


open: probability_density_function.pdf
	open probability_density_function.pdf

o:
	make open

defs: probability_density_function.tex
	defs probability_density_function.tex

spell:
	aspell -c probability_density_function.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile