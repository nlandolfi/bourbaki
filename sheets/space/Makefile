# generated automatically by sheet.go

make:
	make space.pdf

space.tex: sheet.tex
	sheet -mode c -in sheet.tex > space.tex

space.pdf: ../../*.tex ../../trademark.pdf *.tex space.tex
	pdflatex space.tex
	make defs


open: space.pdf
	open space.pdf

o:
	make open

defs: space.tex
	defs space.tex

spell:
	aspell -c space.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile