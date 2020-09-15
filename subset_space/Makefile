subset_space.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex subset_space.tex
	make defs

make:
	make subset_space.pdf

open: subset_space.pdf
	open subset_space.pdf

defs: subset_space.tex
	defs subset_space.tex
