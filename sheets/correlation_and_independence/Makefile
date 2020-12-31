correlation_and_independence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex correlation_and_independence.tex
	make defs

make:
	make correlation_and_independence.pdf

open: correlation_and_independence.pdf
	open correlation_and_independence.pdf

o:
	make open

defs: correlation_and_independence.tex
	defs correlation_and_independence.tex
