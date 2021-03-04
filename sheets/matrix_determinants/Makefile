# generated automatically by sheet.go

make:
	make matrix_determinants.pdf

matrix_determinants.tex: sheet.tex
	sheet -mode c -in sheet.tex > matrix_determinants.tex

matrix_determinants.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_determinants.tex
	pdflatex matrix_determinants.tex
	make defs


open: matrix_determinants.pdf
	open matrix_determinants.pdf

o:
	make open

defs: matrix_determinants.tex
	defs matrix_determinants.tex

spell:
	aspell -c matrix_determinants.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile