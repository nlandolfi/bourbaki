# generated automatically by sheet.go

make:
	make partitions.pdf

partitions.tex: sheet.tex
	sheet -mode c -in sheet.tex > partitions.tex

partitions.pdf: ../../*.tex ../../trademark.pdf *.tex partitions.tex
	pdflatex partitions.tex
	make defs


open: partitions.pdf
	open partitions.pdf

o:
	make open

defs: partitions.tex
	defs partitions.tex

spell:
	aspell -c partitions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile