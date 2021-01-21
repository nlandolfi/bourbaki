cardinality.tex: sheet.tex
	sheet -mode c -in sheet.tex > cardinality.tex

cardinality.pdf: ../../*.tex ../../trademark.pdf *.tex cardinality.tex
	pdflatex cardinality.tex
	make defs

make:
	make cardinality.pdf

open: cardinality.pdf
	open cardinality.pdf

o:
	make open

defs: cardinality.tex
	defs cardinality.tex

spell:
	aspell -c cardinality.tex

s:
	make spell