# generated automatically by sheet.go

make:
	make prelations.pdf

prelations.tex: sheet.tex
	sheet -mode c -in sheet.tex > prelations.tex

prelations.pdf: ../../*.tex ../../trademark.pdf *.tex prelations.tex
	pdflatex prelations.tex
	make defs


open: prelations.pdf
	open prelations.pdf

o:
	make open

defs: prelations.tex
	defs prelations.tex

spell:
	aspell -c prelations.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile