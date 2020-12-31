monotone_sequences.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex monotone_sequences.tex
	make defs

make:
	make monotone_sequences.pdf

open: monotone_sequences.pdf
	open monotone_sequences.pdf

defs: monotone_sequences.tex
	defs monotone_sequences.tex
