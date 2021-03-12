# generated automatically by sheet.go

make:
	make introduction.pdf

introduction.tex: sheet.tex
	sheet -mode c -in sheet.tex > introduction.tex

introduction.pdf: ../../*.tex ../../trademark.pdf *.tex introduction.tex
	pdflatex introduction.tex
	make defs


open: introduction.pdf
	open introduction.pdf

o:
	make open

defs: introduction.tex
	defs introduction.tex

spell:
	aspell -c introduction.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile