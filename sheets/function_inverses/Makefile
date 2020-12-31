function_inverses.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex function_inverses.tex
	make defs

make:
	make function_inverses.pdf

open: function_inverses.pdf
	open function_inverses.pdf

defs: function_inverses.tex
	defs function_inverses.tex
