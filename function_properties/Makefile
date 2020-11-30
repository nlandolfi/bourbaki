function_properties.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex function_properties.tex
	make defs

make:
	make function_properties.pdf

open: function_properties.pdf
	open function_properties.pdf

defs: function_properties.tex
	defs function_properties.tex
