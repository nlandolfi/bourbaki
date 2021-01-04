points.pdf: ../../*.tex ../../trademark.pdf *.tex
	pdflatex points.tex
	make defs

make:
	make points.pdf

open: points.pdf
	open points.pdf

o:
	make open

defs: points.tex
	defs points.tex

spell:
	aspell -c points.tex

s:
	make spell