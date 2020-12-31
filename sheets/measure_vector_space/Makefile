measure_vector_space.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex measure_vector_space.tex
	make defs

make:
	make measure_vector_space.pdf

open: measure_vector_space.pdf
	open measure_vector_space.pdf

o: open
	make open

defs: measure_vector_space.tex
	defs measure_vector_space.tex
