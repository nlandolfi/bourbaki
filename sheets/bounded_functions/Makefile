bounded_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > bounded_functions.tex

bounded_functions.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_functions.tex
	pdflatex bounded_functions.tex
	make defs

make:
	make bounded_functions.pdf

open: bounded_functions.pdf
	open bounded_functions.pdf

o:
	make open

defs: bounded_functions.tex
	defs bounded_functions.tex

spell:
	aspell -c bounded_functions.tex

s:
	make spell