real_limiting_bounds.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_limiting_bounds.tex
	make defs

make:
	make real_limiting_bounds.pdf

open: real_limiting_bounds.pdf
	open real_limiting_bounds.pdf

defs: real_limiting_bounds.tex
	defs real_limiting_bounds.tex
