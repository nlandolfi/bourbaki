real_numbers.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_numbers.tex

real_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex real_numbers.tex
	pdflatex real_numbers.tex
	make defs

make:
	make real_numbers.pdf

open: real_numbers.pdf
	open real_numbers.pdf

o:
	make open

defs: real_numbers.tex
	defs real_numbers.tex

spell:
	aspell -c real_numbers.tex

s:
	make spell