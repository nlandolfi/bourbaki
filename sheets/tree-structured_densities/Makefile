# generated automatically by sheet.go

make:
	make tree-structured_distributions.pdf

tree-structured_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > tree-structured_distributions.tex

tree-structured_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_distributions.tex
	pdflatex tree-structured_distributions.tex
	make defs


open: tree-structured_distributions.pdf
	open tree-structured_distributions.pdf

o:
	make open

defs: tree-structured_distributions.tex
	defs tree-structured_distributions.tex

spell:
	aspell -c tree-structured_distributions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile