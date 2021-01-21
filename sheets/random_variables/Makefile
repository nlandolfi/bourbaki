random_variables.tex: sheet.tex
	sheet -mode c -in sheet.tex > random_variables.tex

random_variables.pdf: ../../*.tex ../../trademark.pdf *.tex random_variables.tex
	pdflatex random_variables.tex
	make defs

make:
	make random_variables.pdf

open: random_variables.pdf
	open random_variables.pdf

o:
	make open

defs: random_variables.tex
	defs random_variables.tex

spell:
	aspell -c random_variables.tex

s:
	make spell