# generated automatically by sheet.go

make:
	make distribution_selection.pdf

distribution_selection.tex: sheet.tex
	sheet -mode c -in sheet.tex > distribution_selection.tex

distribution_selection.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_selection.tex
	pdflatex distribution_selection.tex
	make defs


open: distribution_selection.pdf
	open distribution_selection.pdf

o:
	make open

defs: distribution_selection.tex
	defs distribution_selection.tex

spell:
	aspell -c distribution_selection.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile