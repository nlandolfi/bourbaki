real_series.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_series.tex
	make defs

make:
	make real_series.pdf

open: real_series.pdf
	open real_series.pdf

defs: real_series.tex
	defs real_series.tex
