convex_sets.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex convex_sets.tex
	make defs

make:
	make convex_sets.pdf

open: convex_sets.pdf
	open convex_sets.pdf

defs: convex_sets.tex
	defs convex_sets.tex
