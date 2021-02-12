# generated automatically by sheet.go

make:
	make data_fitting.pdf

data_fitting.tex: sheet.tex
	sheet -mode c -in sheet.tex > data_fitting.tex

data_fitting.pdf: ../../*.tex ../../trademark.pdf *.tex data_fitting.tex
	pdflatex data_fitting.tex
	make defs


open: data_fitting.pdf
	open data_fitting.pdf

o:
	make open

defs: data_fitting.tex
	defs data_fitting.tex

spell:
	aspell -c data_fitting.tex

s:
	make spell