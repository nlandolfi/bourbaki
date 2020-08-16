functions.pdf: ../format.tex ../macros.tex ../trademark.pdf functions.tex macros.tex
	pdflatex functions.tex
	make defs

make:
	make functions.pdf

open: functions.pdf
	open functions.pdf

defs: functions.tex
	defs functions.tex
