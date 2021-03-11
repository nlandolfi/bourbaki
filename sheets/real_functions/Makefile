# generated automatically by sheet.go

make:
	make real_functions.pdf

real_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_functions.tex

real_functions.pdf: ../../*.tex ../../trademark.pdf *.tex real_functions.tex
	pdflatex real_functions.tex
	make defs


open: real_functions.pdf
	open real_functions.pdf

o:
	make open

defs: real_functions.tex
	defs real_functions.tex

spell:
	aspell -c real_functions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile