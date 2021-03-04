# generated automatically by sheet.go

make:
	make tree-structured_density_approximation.pdf

tree-structured_density_approximation.tex: sheet.tex
	sheet -mode c -in sheet.tex > tree-structured_density_approximation.tex

tree-structured_density_approximation.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_density_approximation.tex
	pdflatex tree-structured_density_approximation.tex
	make defs


open: tree-structured_density_approximation.pdf
	open tree-structured_density_approximation.pdf

o:
	make open

defs: tree-structured_density_approximation.tex
	defs tree-structured_density_approximation.tex

spell:
	aspell -c tree-structured_density_approximation.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile