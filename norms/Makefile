norms.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex norms.tex
	make defs

make:
	make norms.pdf

open: norms.pdf
	open norms.pdf

defs: norms.tex
	defs norms.tex
