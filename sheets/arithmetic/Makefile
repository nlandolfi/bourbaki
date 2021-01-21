arithmetic.tex: sheet.tex
	sheet -mode c -in sheet.tex > arithmetic.tex

arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex arithmetic.tex
	pdflatex arithmetic.tex
	make defs

make:
	make arithmetic.pdf

open: arithmetic.pdf
	open arithmetic.pdf

o:
	make open

defs: arithmetic.tex
	defs arithmetic.tex

spell:
	aspell -c arithmetic.tex

s:
	make spell