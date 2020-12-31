function_composition.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex function_composition.tex
	make defs

make:
	make function_composition.pdf

open: function_composition.pdf
	open function_composition.pdf

defs: function_composition.tex
	defs function_composition.tex
