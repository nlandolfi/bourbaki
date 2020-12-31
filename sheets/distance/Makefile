distance.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex distance.tex
	make defs

make:
	make distance.pdf

open: distance.pdf
	open distance.pdf

o:
	make open

defs: distance.tex
	defs distance.tex

spell:
	aspell -c distance.tex

s:
	make spell
