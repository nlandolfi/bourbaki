# generated automatically by sheet.go

make:
	make real_function_graphs.pdf

real_function_graphs.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_function_graphs.tex

real_function_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex real_function_graphs.tex
	pdflatex real_function_graphs.tex
	make defs


open: real_function_graphs.pdf
	open real_function_graphs.pdf

o:
	make open

defs: real_function_graphs.tex
	defs real_function_graphs.tex

spell:
	aspell -c real_function_graphs.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile