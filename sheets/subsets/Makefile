subsets.tex: sheet.tex
	sheet -mode c -in sheet.tex > subsets.tex

subsets.pdf: ../../*.tex ../../trademark.pdf *.tex subsets.tex
	pdflatex subsets.tex
	make defs

make:
	make subsets.pdf

open: subsets.pdf
	open subsets.pdf

o:
	make open

defs: subsets.tex
	defs subsets.tex

spell:
	aspell -c subsets.tex

s:
	make spell