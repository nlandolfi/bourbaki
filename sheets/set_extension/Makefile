set_extension.tex: sheet.tex
	sheet -mode c -in sheet.tex > set_extension.tex

set_extension.pdf: ../../*.tex ../../trademark.pdf *.tex set_extension.tex
	pdflatex set_extension.tex
	make defs

make:
	make set_extension.pdf

open: set_extension.pdf
	open set_extension.pdf

o:
	make open

defs: set_extension.tex
	defs set_extension.tex

spell:
	aspell -c set_extension.tex

s:
	make spell