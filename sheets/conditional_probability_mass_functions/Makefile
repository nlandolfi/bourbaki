# generated automatically by sheet.go

make:
	make conditional_probability_mass_functions.pdf

conditional_probability_mass_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > conditional_probability_mass_functions.tex

conditional_probability_mass_functions.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_probability_mass_functions.tex
	pdflatex conditional_probability_mass_functions.tex
	make defs


open: conditional_probability_mass_functions.pdf
	open conditional_probability_mass_functions.pdf

o:
	make open

defs: conditional_probability_mass_functions.tex
	defs conditional_probability_mass_functions.tex

spell:
	aspell -c conditional_probability_mass_functions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile