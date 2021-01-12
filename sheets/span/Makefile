span.tex: sheet.tex
	sheet -mode c -in sheet.tex > span.tex

span.pdf: ../../*.tex ../../trademark.pdf *.tex span.tex
	pdflatex span.tex
	make defs

make:
	make span.pdf

open: span.pdf
	open span.pdf

o:
	make open

defs: span.tex
	defs span.tex

spell:
	aspell -c span.tex

s:
	make spell