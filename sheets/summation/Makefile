summation.tex: sheet.tex
	sheet -mode c -in sheet.tex > summation.tex

summation.pdf: ../../*.tex ../../trademark.pdf *.tex summation.tex
	pdflatex summation.tex
	make defs

make:
	make summation.pdf

open: summation.pdf
	open summation.pdf

o:
	make open

defs: summation.tex
	defs summation.tex

spell:
	aspell -c summation.tex

s:
	make spell