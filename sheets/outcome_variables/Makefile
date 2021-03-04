# generated automatically by sheet.go

make:
	make outcome_variables.pdf

outcome_variables.tex: sheet.tex
	sheet -mode c -in sheet.tex > outcome_variables.tex

outcome_variables.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_variables.tex
	pdflatex outcome_variables.tex
	make defs


open: outcome_variables.pdf
	open outcome_variables.pdf

o:
	make open

defs: outcome_variables.tex
	defs outcome_variables.tex

spell:
	aspell -c outcome_variables.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile