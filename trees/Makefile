trees.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex trees.tex
	make defs

make:
	make trees.pdf

open: trees.pdf
	open trees.pdf

defs: trees.tex
	defs trees.tex
