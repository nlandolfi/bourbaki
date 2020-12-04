identity_elements.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex identity_elements.tex
	make defs

make:
	make identity_elements.pdf

open: identity_elements.pdf
	open identity_elements.pdf

o:
	make open

defs: identity_elements.tex
	defs identity_elements.tex

spell:
	aspell -c identity_elements.tex

s:
	make spell
