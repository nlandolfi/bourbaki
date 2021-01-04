space.pdf: ../../*.tex ../../trademark.pdf *.tex
	pdflatex space.tex
	make defs

make:
	make space.pdf

open: space.pdf
	open space.pdf

o:
	make open

defs: space.tex
	defs space.tex

spell:
	aspell -c space.tex

s:
	make spell