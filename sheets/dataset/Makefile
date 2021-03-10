# generated automatically by sheet.go

make:
	make dataset.pdf

dataset.tex: sheet.tex
	sheet -mode c -in sheet.tex > dataset.tex

dataset.pdf: ../../*.tex ../../trademark.pdf *.tex dataset.tex
	pdflatex dataset.tex
	make defs


open: dataset.pdf
	open dataset.pdf

o:
	make open

defs: dataset.tex
	defs dataset.tex

spell:
	aspell -c dataset.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile