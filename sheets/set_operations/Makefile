set_operations.tex: sheet.tex
	sheet -mode c -in sheet.tex > set_operations.tex

set_operations.pdf: ../../*.tex ../../trademark.pdf *.tex set_operations.tex
	pdflatex set_operations.tex
	make defs

make:
	make set_operations.pdf

open: set_operations.pdf
	open set_operations.pdf

o:
	make open

defs: set_operations.tex
	defs set_operations.tex

spell:
	aspell -c set_operations.tex

s:
	make spell