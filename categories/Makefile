categories.pdf: ../*.tex *.tex ../trademark.pdf
	pdflatex categories.tex
	make defs

make:
	make categories.pdf

open: categories.pdf
	open categories.pdf

defs: categories.tex
	defs categories.tex
