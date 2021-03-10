# generated automatically by sheet.go

make:
	make real_summation.pdf

real_summation.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_summation.tex

real_summation.pdf: ../../*.tex ../../trademark.pdf *.tex real_summation.tex
	pdflatex real_summation.tex
	make defs


open: real_summation.pdf
	open real_summation.pdf

o:
	make open

defs: real_summation.tex
	defs real_summation.tex

spell:
	aspell -c real_summation.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile