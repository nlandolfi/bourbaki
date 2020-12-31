borel_sigma_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex borel_sigma_algebra.tex
	make defs

make:
	make borel_sigma_algebra.pdf

open: borel_sigma_algebra.pdf
	open borel_sigma_algebra.pdf

defs: borel_sigma_algebra.tex
	defs borel_sigma_algebra.tex
