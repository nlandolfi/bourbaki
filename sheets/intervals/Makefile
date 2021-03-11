# generated automatically by sheet.go

make:
	make intervals.pdf

intervals.tex: sheet.tex
	sheet -mode c -in sheet.tex > intervals.tex

intervals.pdf: ../../*.tex ../../trademark.pdf *.tex intervals.tex
	pdflatex intervals.tex
	make defs


open: intervals.pdf
	open intervals.pdf

o:
	make open

defs: intervals.tex
	defs intervals.tex

spell:
	aspell -c intervals.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile