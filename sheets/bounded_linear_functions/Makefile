bounded_linear_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > bounded_linear_functions.tex

bounded_linear_functions.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_linear_functions.tex
	pdflatex bounded_linear_functions.tex
	make defs

make:
	make bounded_linear_functions.pdf

open: bounded_linear_functions.pdf
	open bounded_linear_functions.pdf

o:
	make open

defs: bounded_linear_functions.tex
	defs bounded_linear_functions.tex

spell:
	aspell -c bounded_linear_functions.tex

s:
	make spell