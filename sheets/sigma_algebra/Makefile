sigma_algebra.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex sigma_algebra.tex
	make defs

make:
	make sigma_algebra.pdf

open: sigma_algebra.pdf
	open sigma_algebra.pdf

defs: sigma_algebra.tex
	defs sigma_algebra.tex
