pushforward_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex pushforward_measures.tex
	make defs

make:
	make pushforward_measures.pdf

open: pushforward_measures.pdf
	open pushforward_measures.pdf

defs: pushforward_measures.tex
	defs pushforward_measures.tex
