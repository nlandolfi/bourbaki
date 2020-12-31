signed_set_decomposition.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex signed_set_decomposition.tex
	make defs

make:
	make signed_set_decomposition.pdf

open: signed_set_decomposition.pdf
	open signed_set_decomposition.pdf

o:
	make open

defs: signed_set_decomposition.tex
	defs signed_set_decomposition.tex
