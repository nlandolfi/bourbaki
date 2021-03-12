# generated automatically by sheet.go

make:
	make differentiability.pdf

differentiability.tex: sheet.tex
	sheet -mode c -in sheet.tex > differentiability.tex

differentiability.pdf: ../../*.tex ../../trademark.pdf *.tex differentiability.tex
	pdflatex differentiability.tex
	make defs


open: differentiability.pdf
	open differentiability.pdf

o:
	make open

defs: differentiability.tex
	defs differentiability.tex

spell:
	aspell -c differentiability.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile