interval_length.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex interval_length.tex
	make defs

make:
	make interval_length.pdf

open: interval_length.pdf
	open interval_length.pdf

o:
	make open

defs: interval_length.tex
	defs interval_length.tex
