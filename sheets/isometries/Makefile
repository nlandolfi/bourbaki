isometries.tex: sheet.tex
	sheet -mode c -in sheet.tex > isometries.tex

isometries.pdf: ../../*.tex ../../trademark.pdf *.tex isometries.tex
	pdflatex isometries.tex
	make defs

make:
	make isometries.pdf

open: isometries.pdf
	open isometries.pdf

o:
	make open

defs: isometries.tex
	defs isometries.tex

spell:
	aspell -c isometries.tex

s:
	make spell