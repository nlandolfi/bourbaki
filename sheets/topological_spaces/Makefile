topological_spaces.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex topological_spaces.tex
	make defs

make:
	make topological_spaces.pdf

open: topological_spaces.pdf
	open topological_spaces.pdf

defs: topological_spaces.tex
	defs topological_spaces.tex
