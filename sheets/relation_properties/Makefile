# generated automatically by sheet.go

make:
	make relation_properties.pdf

relation_properties.tex: sheet.tex
	sheet -mode c -in sheet.tex > relation_properties.tex

relation_properties.pdf: ../../*.tex ../../trademark.pdf *.tex relation_properties.tex
	pdflatex relation_properties.tex
	make defs


open: relation_properties.pdf
	open relation_properties.pdf

o:
	make open

defs: relation_properties.tex
	defs relation_properties.tex

spell:
	aspell -c relation_properties.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile