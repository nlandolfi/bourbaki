# generated automatically by sheet.go

make:
	make trees.pdf

trees.tex: sheet.tex
	sheet -mode c -in sheet.tex > trees.tex

trees.pdf: ../../*.tex ../../trademark.pdf *.tex trees.tex
	pdflatex trees.tex
	make defs


open: trees.pdf
	open trees.pdf

o:
	make open

defs: trees.tex
	defs trees.tex

spell:
	aspell -c trees.tex

s:
	make spell