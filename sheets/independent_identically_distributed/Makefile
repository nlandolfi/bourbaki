independent_identically_distributed.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex independent_identically_distributed.tex
	make defs

make:
	make independent_identically_distributed.pdf

open: independent_identically_distributed.pdf
	open independent_identically_distributed.pdf

o:
	make open

defs: independent_identically_distributed.tex
	defs independent_identically_distributed.tex
