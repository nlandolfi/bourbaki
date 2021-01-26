sentences.tex: sheet.tex
	sheet -mode c -in sheet.tex > sentences.tex

sentences.pdf: ../../*.tex ../../trademark.pdf *.tex sentences.tex
	pdflatex sentences.tex
	make defs

make:
	make sentences.pdf

open: sentences.pdf
	open sentences.pdf

o:
	make open

defs: sentences.tex
	defs sentences.tex

spell:
	aspell -c sentences.tex

s:
	make spell