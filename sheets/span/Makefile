# generated automatically by sheet.go

make:
	make span.pdf

span.tex: sheet.tex
	sheet -mode c -in sheet.tex > span.tex

span.pdf: ../../*.tex ../../trademark.pdf *.tex span.tex
	pdflatex span.tex
	make defs


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

remake:
	sheet -mode mk > Makefile