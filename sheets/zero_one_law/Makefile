zero_one_law.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex zero_one_law.tex
	make defs

make:
	make zero_one_law.pdf

open: zero_one_law.pdf
	open zero_one_law.pdf

o:
	make open

defs: zero_one_law.tex
	defs zero_one_law.tex
