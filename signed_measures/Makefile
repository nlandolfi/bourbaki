signed_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex signed_measures.tex
	make defs

make:
	make signed_measures.pdf

open: signed_measures.pdf
	open signed_measures.pdf

o: open
	make open

defs: signed_measures.tex
	defs signed_measures.tex
