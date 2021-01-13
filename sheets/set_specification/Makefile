set_specification.tex: sheet.tex
	sheet -mode c -in sheet.tex > set_specification.tex

set_specification.pdf: ../../*.tex ../../trademark.pdf *.tex set_specification.tex
	pdflatex set_specification.tex
	make defs

make:
	make set_specification.pdf

open: set_specification.pdf
	open set_specification.pdf

o:
	make open

defs: set_specification.tex
	defs set_specification.tex

spell:
	aspell -c set_specification.tex

s:
	make spell