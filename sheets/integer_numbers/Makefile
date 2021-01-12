integer_numbers.tex: sheet.tex
	sheet -mode c -in sheet.tex > integer_numbers.tex

integer_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex integer_numbers.tex
	pdflatex integer_numbers.tex
	make defs

make:
	make integer_numbers.pdf

open: integer_numbers.pdf
	open integer_numbers.pdf

o:
	make open

defs: integer_numbers.tex
	defs integer_numbers.tex

spell:
	aspell -c integer_numbers.tex

s:
	make spell