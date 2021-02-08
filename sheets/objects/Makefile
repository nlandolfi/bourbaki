objects.tex: sheet.tex
	sheet -mode c -in sheet.tex > objects.tex

objects.pdf: ../../*.tex ../../trademark.pdf *.tex objects.tex
	pdflatex objects.tex
	make defs

make:
	make objects.pdf

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