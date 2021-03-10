# generated automatically by sheet.go

make:
	make almost_everywhere.pdf

almost_everywhere.tex: sheet.tex
	sheet -mode c -in sheet.tex > almost_everywhere.tex

almost_everywhere.pdf: ../../*.tex ../../trademark.pdf *.tex almost_everywhere.tex
	pdflatex almost_everywhere.tex
	make defs


open: almost_everywhere.pdf
	open almost_everywhere.pdf

o:
	make open

defs: almost_everywhere.tex
	defs almost_everywhere.tex

spell:
	aspell -c almost_everywhere.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile