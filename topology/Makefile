topology.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex topology.tex
	make defs

make:
	make topology.pdf

open: topology.pdf
	open topology.pdf

defs: topology.tex
	defs topology.tex
