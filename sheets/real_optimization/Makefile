# generated automatically by sheet.go

make:
	make real_optimization.pdf

real_optimization.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_optimization.tex

real_optimization.pdf: ../../*.tex ../../trademark.pdf *.tex real_optimization.tex
	pdflatex real_optimization.tex
	make defs


open: real_optimization.pdf
	open real_optimization.pdf

o:
	make open

defs: real_optimization.tex
	defs real_optimization.tex

spell:
	aspell -c real_optimization.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile