convex_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex convex_functions.tex
	make defs

make:
	make convex_functions.pdf

open: convex_functions.pdf
	open convex_functions.pdf

o:
	make open

defs: convex_functions.tex
	defs convex_functions.tex
