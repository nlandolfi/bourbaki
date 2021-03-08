# generated automatically by sheet.go

make:
	make relations.pdf

relations.tex: sheet.tex
	sheet -mode c -in sheet.tex > relations.tex

relations.pdf: ../../*.tex ../../trademark.pdf *.tex relations.tex
	pdflatex relations.tex
	make defs


open: relations.pdf
	open relations.pdf

o:
	make open

defs: relations.tex
	defs relations.tex

spell:
	aspell -c relations.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile