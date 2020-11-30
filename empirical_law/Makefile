empirical_law.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex empirical_law.tex
	make defs

make:
	make empirical_law.pdf

open: empirical_law.pdf
	open empirical_law.pdf

o:
	make open

defs: empirical_law.tex
	defs empirical_law.tex

spell:
	aspell -c empirical_law.tex

s:
	make spell
