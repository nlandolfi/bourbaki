set_examples.tex: sheet.tex
	sheet -mode c -in sheet.tex > set_examples.tex

set_examples.pdf: ../../*.tex ../../trademark.pdf *.tex set_examples.tex
	pdflatex set_examples.tex
	make defs

make:
	make set_examples.pdf

open: set_examples.pdf
	open set_examples.pdf

o:
	make open

defs: set_examples.tex
	defs set_examples.tex

spell:
	aspell -c set_examples.tex

s:
	make spell