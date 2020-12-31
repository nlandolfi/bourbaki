pointwise_vs_measure_limits.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex pointwise_vs_measure_limits.tex
	make defs

make:
	make pointwise_vs_measure_limits.pdf

open: pointwise_vs_measure_limits.pdf
	open pointwise_vs_measure_limits.pdf

o:
	make open

defs: pointwise_vs_measure_limits.tex
	defs pointwise_vs_measure_limits.tex
