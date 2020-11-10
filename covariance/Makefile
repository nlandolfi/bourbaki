covariance.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex covariance.tex
	make defs

make:
	make covariance.pdf

open: covariance.pdf
	open covariance.pdf

o:
	make open

defs: covariance.tex
	defs covariance.tex
