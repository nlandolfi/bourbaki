ordered_pairs.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex ordered_pairs.tex
	make defs

make:
	make ordered_pairs.pdf

open: ordered_pairs.pdf
	open ordered_pairs.pdf

defs: ordered_pairs.tex
	defs ordered_pairs.tex
