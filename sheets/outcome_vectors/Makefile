# generated automatically by sheet.go

make:
	make outcome_vectors.pdf

outcome_vectors.tex: sheet.tex
	sheet -mode c -in sheet.tex > outcome_vectors.tex

outcome_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_vectors.tex
	pdflatex outcome_vectors.tex
	make defs


open: outcome_vectors.pdf
	open outcome_vectors.pdf

o:
	make open

defs: outcome_vectors.tex
	defs outcome_vectors.tex

spell:
	aspell -c outcome_vectors.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile