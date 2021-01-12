length.tex: sheet.tex
	sheet -mode c -in sheet.tex > length.tex

length.pdf: ../../*.tex ../../trademark.pdf *.tex length.tex
	pdflatex length.tex
	make defs

make:
	make length.pdf

open: length.pdf
	open length.pdf

o:
	make open

defs: length.tex
	defs length.tex

spell:
	aspell -c length.tex

s:
	make spell