# generated automatically by sheet.go

make:
	make evaluators.pdf

evaluators.tex: sheet.tex
	sheet -mode c -in sheet.tex > evaluators.tex

evaluators.pdf: ../../*.tex ../../trademark.pdf *.tex evaluators.tex
	pdflatex evaluators.tex
	make defs


open: evaluators.pdf
	open evaluators.pdf

o:
	make open

defs: evaluators.tex
	defs evaluators.tex

spell:
	aspell -c evaluators.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile