# generated automatically by sheet.go

make:
	make objects.pdf

objects.tex: sheet.tex
	sheet -mode c -in sheet.tex > objects.tex

objects.pdf: ../../*.tex ../../trademark.pdf *.tex objects.tex
	pdflatex objects.tex
	make defs


open: objects.pdf
	open objects.pdf

o:
	make open

defs: objects.tex
	defs objects.tex

spell:
	aspell -c objects.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile