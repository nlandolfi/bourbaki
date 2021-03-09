# generated automatically by sheet.go

make:
	make tree-structured_densities.pdf

tree-structured_densities.tex: sheet.tex
	sheet -mode c -in sheet.tex > tree-structured_densities.tex

tree-structured_densities.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_densities.tex
	pdflatex tree-structured_densities.tex
	make defs


open: tree-structured_densities.pdf
	open tree-structured_densities.pdf

o:
	make open

defs: tree-structured_densities.tex
	defs tree-structured_densities.tex

spell:
	aspell -c tree-structured_densities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile