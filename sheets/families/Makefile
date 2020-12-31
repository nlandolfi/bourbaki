families.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex families.tex
	make defs

make:
	make families.pdf

open: families.pdf
	open families.pdf

defs: families.tex
	defs families.tex
