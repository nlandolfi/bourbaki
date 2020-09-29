sequences.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex sequences.tex
	make defs

make:
	make sequences.pdf

open: sequences.pdf
	open sequences.pdf

defs: sequences.tex
	defs sequences.tex
