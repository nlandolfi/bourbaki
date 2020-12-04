operation_properties.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex operation_properties.tex
	make defs

make:
	make operation_properties.pdf

open: operation_properties.pdf
	open operation_properties.pdf

defs: operation_properties.tex
	defs operation_properties.tex

o:
	make open
