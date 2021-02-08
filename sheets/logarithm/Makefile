# generated automatically by sheet.go

make:
	make logarithm.pdf

logarithm.tex: sheet.tex
	sheet -mode c -in sheet.tex > logarithm.tex

logarithm.pdf: ../../*.tex ../../trademark.pdf *.tex logarithm.tex
	pdflatex logarithm.tex
	make defs


open: logarithm.pdf
	open logarithm.pdf

o:
	make open

defs: logarithm.tex
	defs logarithm.tex

spell:
	aspell -c logarithm.tex

s:
	make spell