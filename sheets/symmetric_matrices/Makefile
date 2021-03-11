# generated automatically by sheet.go

make:
	make symmetric_matrices.pdf

symmetric_matrices.tex: sheet.tex
	sheet -mode c -in sheet.tex > symmetric_matrices.tex

symmetric_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_matrices.tex
	pdflatex symmetric_matrices.tex
	make defs


open: symmetric_matrices.pdf
	open symmetric_matrices.pdf

o:
	make open

defs: symmetric_matrices.tex
	defs symmetric_matrices.tex

spell:
	aspell -c symmetric_matrices.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile