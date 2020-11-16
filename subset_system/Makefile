subset_system.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex subset_system.tex
	make defs

make:
	make subset_system.pdf

open: subset_system.pdf
	open subset_system.pdf

o:
	make open

defs: subset_system.tex
	defs subset_system.tex
