distance_asymmetry.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex distance_asymmetry.tex
	make defs

make:
	make distance_asymmetry.pdf

open: distance_asymmetry.pdf
	open distance_asymmetry.pdf

o:
	make open

defs: distance_asymmetry.tex
	defs distance_asymmetry.tex

spell:
	aspell -c distance_asymmetry.tex

s:
	make spell
