algebras.tex: sheet.tex
	sheet -mode c -in sheet.tex > algebras.tex

algebras.pdf: ../../*.tex ../../trademark.pdf *.tex algebras.tex
	pdflatex algebras.tex
	make defs

make:
	make algebras.pdf

open: algebras.pdf
	open algebras.pdf

o:
	make open

defs: algebras.tex
	defs algebras.tex

spell:
	aspell -c algebras.tex

s:
	make spell