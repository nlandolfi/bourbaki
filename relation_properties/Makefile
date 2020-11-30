relation_properties.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex relation_properties.tex
	make defs

make:
	make relation_properties.pdf

open: relation_properties.pdf
	open relation_properties.pdf

defs: relation_properties.tex
	defs relation_properties.tex
