subset_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex subset_algebra.tex
	make defs

make:
	make subset_algebra.pdf

open: subset_algebra.pdf
	open subset_algebra.pdf

o:
	make open

defs: subset_algebra.tex
	defs subset_algebra.tex
