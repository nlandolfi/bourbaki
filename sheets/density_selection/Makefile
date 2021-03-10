# generated automatically by sheet.go

make:
	make density_selection.pdf

density_selection.tex: sheet.tex
	sheet -mode c -in sheet.tex > density_selection.tex

density_selection.pdf: ../../*.tex ../../trademark.pdf *.tex density_selection.tex
	pdflatex density_selection.tex
	make defs


open: density_selection.pdf
	open density_selection.pdf

o:
	make open

defs: density_selection.tex
	defs density_selection.tex

spell:
	aspell -c density_selection.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile