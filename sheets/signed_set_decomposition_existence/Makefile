signed_set_decomposition_existence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex signed_set_decomposition_existence.tex
	make defs

make:
	make signed_set_decomposition_existence.pdf

open: signed_set_decomposition_existence.pdf
	open signed_set_decomposition_existence.pdf

o:
	make open

defs: signed_set_decomposition_existence.tex
	defs signed_set_decomposition_existence.tex
