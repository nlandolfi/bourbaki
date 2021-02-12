# generated automatically by sheet.go

make:
	make matrices.pdf

matrices.tex: sheet.tex
	sheet -mode c -in sheet.tex > matrices.tex

matrices.pdf: ../../*.tex ../../trademark.pdf *.tex matrices.tex
	pdflatex matrices.tex
	make defs


open: matrices.pdf
	open matrices.pdf

o:
	make open

defs: matrices.tex
	defs matrices.tex

spell:
	aspell -c matrices.tex

s:
	make spell