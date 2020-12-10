solving_equations.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex solving_equations.tex
	make defs

make:
	make solving_equations.pdf

open: solving_equations.pdf
	open solving_equations.pdf

defs: solving_equations.tex
	defs solving_equations.tex

o:
	make open
