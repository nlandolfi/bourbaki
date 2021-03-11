# generated automatically by sheet.go

make:
	make transpose.pdf

transpose.tex: sheet.tex
	sheet -mode c -in sheet.tex > transpose.tex

transpose.pdf: ../../*.tex ../../trademark.pdf *.tex transpose.tex
	pdflatex transpose.tex
	make defs


open: transpose.pdf
	open transpose.pdf

o:
	make open

defs: transpose.tex
	defs transpose.tex

spell:
	aspell -c transpose.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile