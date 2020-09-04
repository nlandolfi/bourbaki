graphs.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex graphs.tex
	make defs

make:
	make graphs.pdf

open: graphs.pdf
	open graphs.pdf

defs: graphs.tex
	defs graphs.tex
