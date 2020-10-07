characteristic_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex characteristic_functions.tex
	make defs

make:
	make characteristic_functions.pdf

open: characteristic_functions.pdf
	open characteristic_functions.pdf

defs: characteristic_functions.tex
	defs characteristic_functions.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
