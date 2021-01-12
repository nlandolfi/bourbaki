norms.tex: sheet.tex
	sheet -mode c -in sheet.tex > norms.tex

norms.pdf: ../../*.tex ../../trademark.pdf *.tex norms.tex
	pdflatex norms.tex
	make defs

make:
	make norms.pdf

open: norms.pdf
	open norms.pdf

o:
	make open

defs: norms.tex
	defs norms.tex

spell:
	aspell -c norms.tex

s:
	make spell