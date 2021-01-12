fields.tex: sheet.tex
	sheet -mode c -in sheet.tex > fields.tex

fields.pdf: ../../*.tex ../../trademark.pdf *.tex fields.tex
	pdflatex fields.tex
	make defs

make:
	make fields.pdf

open: fields.pdf
	open fields.pdf

o:
	make open

defs: fields.tex
	defs fields.tex

spell:
	aspell -c fields.tex

s:
	make spell