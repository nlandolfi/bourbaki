linear_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > linear_functions.tex

linear_functions.pdf: ../../*.tex ../../trademark.pdf *.tex linear_functions.tex
	pdflatex linear_functions.tex
	make defs

make:
	make linear_functions.pdf

open: linear_functions.pdf
	open linear_functions.pdf

o:
	make open

defs: linear_functions.tex
	defs linear_functions.tex

spell:
	aspell -c linear_functions.tex

s:
	make spell