vectors.pdf: ../../*.tex ../../trademark.pdf *.tex
	pdflatex vectors.tex
	make defs

make:
	make vectors.pdf

open: vectors.pdf
	open vectors.pdf

defs: vectors.tex
	defs vectors.tex

o:
	make open
