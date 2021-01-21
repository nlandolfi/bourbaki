equations.tex: sheet.tex
	sheet -mode c -in sheet.tex > equations.tex

equations.pdf: ../../*.tex ../../trademark.pdf *.tex equations.tex
	pdflatex equations.tex
	make defs

make:
	make equations.pdf

open: equations.pdf
	open equations.pdf

o:
	make open

defs: equations.tex
	defs equations.tex

spell:
	aspell -c equations.tex

s:
	make spell