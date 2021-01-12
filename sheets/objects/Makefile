objects.pdf: ../../*.tex ../../trademark.pdf *.tex
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