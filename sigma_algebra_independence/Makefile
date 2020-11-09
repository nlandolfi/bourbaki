sigma_algebra_independence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex sigma_algebra_independence.tex
	make defs

make:
	make sigma_algebra_independence.pdf

open: sigma_algebra_independence.pdf
	open sigma_algebra_independence.pdf

o:
	make open

defs: sigma_algebra_independence.tex
	defs sigma_algebra_independence.tex
