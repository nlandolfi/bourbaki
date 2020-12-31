fields.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex fields.tex
	make defs

make:
	make fields.pdf

open: fields.pdf
	open fields.pdf

defs: fields.tex
	defs fields.tex
