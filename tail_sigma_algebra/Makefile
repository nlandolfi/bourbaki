tail_sigma_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex tail_sigma_algebra.tex
	make defs

make:
	make tail_sigma_algebra.pdf

open: tail_sigma_algebra.pdf
	open tail_sigma_algebra.pdf

o:
	make open

defs: tail_sigma_algebra.tex
	defs tail_sigma_algebra.tex
