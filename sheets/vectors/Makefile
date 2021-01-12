vectors.tex: sheet.tex
	sheet -mode c -in sheet.tex > vectors.tex

vectors.pdf: ../../*.tex ../../trademark.pdf *.tex vectors.tex
	pdflatex vectors.tex
	make defs

make:
	make vectors.pdf

open: vectors.pdf
	open vectors.pdf

o:
	make open

defs: vectors.tex
	defs vectors.tex

spell:
	aspell -c vectors.tex

s:
	make spell