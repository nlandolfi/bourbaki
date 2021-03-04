# generated automatically by sheet.go

make:
	make records.pdf

records.tex: sheet.tex
	sheet -mode c -in sheet.tex > records.tex

records.pdf: ../../*.tex ../../trademark.pdf *.tex records.tex
	pdflatex records.tex
	make defs


open: records.pdf
	open records.pdf

o:
	make open

defs: records.tex
	defs records.tex

spell:
	aspell -c records.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile