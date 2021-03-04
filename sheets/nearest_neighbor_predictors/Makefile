# generated automatically by sheet.go

make:
	make nearest_neighbor_predictors.pdf

nearest_neighbor_predictors.tex: sheet.tex
	sheet -mode c -in sheet.tex > nearest_neighbor_predictors.tex

nearest_neighbor_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex nearest_neighbor_predictors.tex
	pdflatex nearest_neighbor_predictors.tex
	make defs


open: nearest_neighbor_predictors.pdf
	open nearest_neighbor_predictors.pdf

o:
	make open

defs: nearest_neighbor_predictors.tex
	defs nearest_neighbor_predictors.tex

spell:
	aspell -c nearest_neighbor_predictors.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile