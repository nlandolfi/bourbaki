functionals.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex functionals.tex
	make defs

make:
	make functionals.pdf

open: functionals.pdf
	open functionals.pdf

defs: functionals.tex
	defs functionals.tex
