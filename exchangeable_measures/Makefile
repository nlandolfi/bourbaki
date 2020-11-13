exchangeable_measures.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex exchangeable_measures.tex
	make defs

make:
	make exchangeable_measures.pdf

open: exchangeable_measures.pdf
	open exchangeable_measures.pdf

o:
	make open

defs: exchangeable_measures.tex
	defs exchangeable_measures.tex
