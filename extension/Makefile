extension.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex extension.tex
	make defs

make:
	make extension.pdf

open: extension.pdf
	open extension.pdf

o:
	make open

defs: extension.tex
	defs extension.tex

spell:
	aspell -c extension.tex

s:
	make spell
