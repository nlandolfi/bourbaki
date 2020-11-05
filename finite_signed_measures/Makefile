finite_signed_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex finite_signed_measures.tex
	make defs

make:
	make finite_signed_measures.pdf

open: finite_signed_measures.pdf
	open finite_signed_measures.pdf

o: open
	make open

defs: finite_signed_measures.tex
	defs finite_signed_measures.tex
