topological_space.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex topological_space.tex
	make defs

make:
	make topological_space.pdf

open: topological_space.pdf
	open topological_space.pdf

defs: topological_space.tex
	defs topological_space.tex
