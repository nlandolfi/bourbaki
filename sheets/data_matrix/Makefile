# generated automatically by sheet.go

make:
	make data_matrix.pdf

data_matrix.tex: sheet.tex
	sheet -mode c -in sheet.tex > data_matrix.tex

data_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex data_matrix.tex
	pdflatex data_matrix.tex
	make defs


open: data_matrix.pdf
	open data_matrix.pdf

o:
	make open

defs: data_matrix.tex
	defs data_matrix.tex

spell:
	aspell -c data_matrix.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile