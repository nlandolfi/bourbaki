real_intervals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_intervals.tex
	make defs

make:
	make real_intervals.pdf

open: real_intervals.pdf
	open real_intervals.pdf

defs: real_intervals.tex
	defs real_intervals.tex
