# generated automatically by sheet.go

make:
	make entropy.pdf

entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > entropy.tex

entropy.pdf: ../../*.tex ../../trademark.pdf *.tex entropy.tex
	pdflatex entropy.tex
	make defs


open: entropy.pdf
	open entropy.pdf

o:
	make open

defs: entropy.tex
	defs entropy.tex

spell:
	aspell -c entropy.tex

s:
	make spell