complex_numbers.tex: sheet.tex
	sheet -mode c -in sheet.tex > complex_numbers.tex

complex_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex complex_numbers.tex
	pdflatex complex_numbers.tex
	make defs

make:
	make complex_numbers.pdf

open: complex_numbers.pdf
	open complex_numbers.pdf

o:
	make open

defs: complex_numbers.tex
	defs complex_numbers.tex

spell:
	aspell -c complex_numbers.tex

s:
	make spell