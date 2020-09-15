real_limits.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_limits.tex
	make defs

make:
	make real_limits.pdf

open: real_limits.pdf
	open real_limits.pdf

defs: real_limits.tex
	defs real_limits.tex
