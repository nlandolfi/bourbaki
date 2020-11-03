functions.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex functions.tex
	make defs

make:
	make functions.pdf

open: functions.pdf
	open functions.pdf

o: open
	make open

defs: functions.tex
	defs functions.tex

clean:
	rm *.aux
	rm *.out
	rm *.log
