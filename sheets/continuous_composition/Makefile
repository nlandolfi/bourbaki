continuous_composition.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex continuous_composition.tex
	make defs

make:
	make continuous_composition.pdf

open: continuous_composition.pdf
	open continuous_composition.pdf

o:
	make open

defs: continuous_composition.tex
	defs continuous_composition.tex

spell:
	aspell -c continuous_composition.tex

s:
	make spell
