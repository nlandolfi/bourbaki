groups.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex groups.tex
	make defs

make:
	make groups.pdf

open: groups.pdf
	open groups.pdf

defs: groups.tex
	defs groups.tex
