matroids.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex matroids.tex
	make defs

make:
	make matroids.pdf

open: matroids.pdf
	open matroids.pdf

defs: matroids.tex
	defs matroids.tex
