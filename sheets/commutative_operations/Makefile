commutative_operations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex commutative_operations.tex
	make defs

make:
	make commutative_operations.pdf

open: commutative_operations.pdf
	open commutative_operations.pdf

defs: commutative_operations.tex
	defs commutative_operations.tex

o:
	make open
