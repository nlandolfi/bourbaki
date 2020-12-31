generated_sigma_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex generated_sigma_algebra.tex
	make defs

make:
	make generated_sigma_algebra.pdf

open: generated_sigma_algebra.pdf
	open generated_sigma_algebra.pdf

defs: generated_sigma_algebra.tex
	defs generated_sigma_algebra.tex
