supremum_norm.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex supremum_norm.tex
	make defs

make:
	make supremum_norm.pdf

open: supremum_norm.pdf
	open supremum_norm.pdf

defs: supremum_norm.tex
	defs supremum_norm.tex
