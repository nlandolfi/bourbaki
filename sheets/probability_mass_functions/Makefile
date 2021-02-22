# generated automatically by sheet.go

make:
	make probability_mass_functions.pdf

probability_mass_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_mass_functions.tex

probability_mass_functions.pdf: ../../*.tex ../../trademark.pdf *.tex probability_mass_functions.tex
	pdflatex probability_mass_functions.tex
	make defs


open: probability_mass_functions.pdf
	open probability_mass_functions.pdf

o:
	make open

defs: probability_mass_functions.tex
	defs probability_mass_functions.tex

spell:
	aspell -c probability_mass_functions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile