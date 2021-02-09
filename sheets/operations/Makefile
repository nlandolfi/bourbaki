# generated automatically by sheet.go

make:
	make operations.pdf

operations.tex: sheet.tex
	sheet -mode c -in sheet.tex > operations.tex

operations.pdf: ../../*.tex ../../trademark.pdf *.tex operations.tex
	pdflatex operations.tex
	make defs


open: operations.pdf
	open operations.pdf

o:
	make open

defs: operations.tex
	defs operations.tex

spell:
	aspell -c operations.tex

s:
	make spell