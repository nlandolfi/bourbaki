relations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex relations.tex
	make defs

make:
	make relations.pdf

open: relations.pdf
	open relations.pdf

defs: relations.tex
	defs relations.tex
