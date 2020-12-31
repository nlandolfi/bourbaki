tail_measure_lower_bound.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex tail_measure_lower_bound.tex
	make defs

make:
	make tail_measure_lower_bound.pdf

open: tail_measure_lower_bound.pdf
	open tail_measure_lower_bound.pdf

o:
	make open

defs: tail_measure_lower_bound.tex
	defs tail_measure_lower_bound.tex
