measure_properties.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measure_properties.tex
	make defs

make:
	make measure_properties.pdf

open: measure_properties.pdf
	open measure_properties.pdf

defs: measure_properties.tex
	defs measure_properties.tex
