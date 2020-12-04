algebras.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex algebras.tex
	make defs

make:
	make algebras.pdf

open: algebras.pdf
	open algebras.pdf

defs: algebras.tex
	defs algebras.tex

o:
	make open
