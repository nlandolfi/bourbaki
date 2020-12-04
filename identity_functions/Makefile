identity_functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex identity_functions.tex
	make defs

make:
	make identity_functions.pdf

open: identity_functions.pdf
	open identity_functions.pdf

o: open
	make open

defs: identity_functions.tex
	defs identity_functions.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
