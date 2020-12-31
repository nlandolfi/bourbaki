signed_measure_examples.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex signed_measure_examples.tex
	make defs

make:
	make signed_measure_examples.pdf

open: signed_measure_examples.pdf
	open signed_measure_examples.pdf

defs: signed_measure_examples.tex
	defs signed_measure_examples.tex
