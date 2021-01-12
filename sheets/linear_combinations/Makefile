linear_combinations.tex: sheet.tex
	sheet -mode c -in sheet.tex > linear_combinations.tex

linear_combinations.pdf: ../../*.tex ../../trademark.pdf *.tex linear_combinations.tex
	pdflatex linear_combinations.tex
	make defs

make:
	make linear_combinations.pdf

open: linear_combinations.pdf
	open linear_combinations.pdf

o:
	make open

defs: linear_combinations.tex
	defs linear_combinations.tex

spell:
	aspell -c linear_combinations.tex

s:
	make spell