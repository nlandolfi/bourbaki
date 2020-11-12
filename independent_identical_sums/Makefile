independent_identical_sums.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex independent_identical_sums.tex
	make defs

make:
	make independent_identical_sums.pdf

open: independent_identical_sums.pdf
	open independent_identical_sums.pdf

o:
	make open

defs: independent_identical_sums.tex
	defs independent_identical_sums.tex
