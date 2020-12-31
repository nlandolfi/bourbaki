real_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex real_functions.tex
	make defs

make:
	make real_functions.pdf

open: real_functions.pdf
	open real_functions.pdf

defs: real_functions.tex
	defs real_functions.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
