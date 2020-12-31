homomorphism.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex homomorphism.tex
	make defs

make:
	make homomorphism.pdf

open: homomorphism.pdf
	open homomorphism.pdf

defs: homomorphism.tex
	defs homomorphism.tex
