metrics.tex: sheet.tex
	sheet -mode c -in sheet.tex > metrics.tex

metrics.pdf: ../../*.tex ../../trademark.pdf *.tex metrics.tex
	pdflatex metrics.tex
	make defs

make:
	make metrics.pdf

open: metrics.pdf
	open metrics.pdf

o:
	make open

defs: metrics.tex
	defs metrics.tex

spell:
	aspell -c metrics.tex

s:
	make spell