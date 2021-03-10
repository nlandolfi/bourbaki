# generated automatically by sheet.go

make:
	make real_series.pdf

real_series.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_series.tex

real_series.pdf: ../../*.tex ../../trademark.pdf *.tex real_series.tex
	pdflatex real_series.tex
	make defs


open: real_series.pdf
	open real_series.pdf

o:
	make open

defs: real_series.tex
	defs real_series.tex

spell:
	aspell -c real_series.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile