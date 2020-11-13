exchangeable_singular_decomposition.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex exchangeable_singular_decomposition.tex
	make defs

make:
	make exchangeable_singular_decomposition.pdf

open: exchangeable_singular_decomposition.pdf
	open exchangeable_singular_decomposition.pdf

o:
	make open

defs: exchangeable_singular_decomposition.tex
	defs exchangeable_singular_decomposition.tex
