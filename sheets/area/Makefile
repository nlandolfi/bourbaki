area.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex area.tex
	make defs

make:
	make area.pdf

open: area.pdf
	open area.pdf

defs: area.tex
	defs area.tex
