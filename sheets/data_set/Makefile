# generated automatically by sheet.go

make:
	make data_set.pdf

data_set.tex: sheet.tex
	sheet -mode c -in sheet.tex > data_set.tex

data_set.pdf: ../../*.tex ../../trademark.pdf *.tex data_set.tex
	pdflatex data_set.tex
	make defs


open: data_set.pdf
	open data_set.pdf

o:
	make open

defs: data_set.tex
	defs data_set.tex

spell:
	aspell -c data_set.tex

s:
	make spell