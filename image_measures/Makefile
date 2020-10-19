image_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex image_measures.tex
	make defs

make:
	make image_measures.pdf

open: image_measures.pdf
	open image_measures.pdf

defs: image_measures.tex
	defs image_measures.tex
