# generated automatically by sheet.go

make:
	make predictors.pdf

predictors.tex: sheet.tex
	sheet -mode c -in sheet.tex > predictors.tex

predictors.pdf: ../../*.tex ../../trademark.pdf *.tex predictors.tex
	pdflatex predictors.tex
	make defs


open: predictors.pdf
	open predictors.pdf

o:
	make open

defs: predictors.tex
	defs predictors.tex

spell:
	aspell -c predictors.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile