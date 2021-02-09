# generated automatically by sheet.go

make:
	make cross_entropy.pdf

cross_entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > cross_entropy.tex

cross_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex cross_entropy.tex
	pdflatex cross_entropy.tex
	make defs


open: cross_entropy.pdf
	open cross_entropy.pdf

o:
	make open

defs: cross_entropy.tex
	defs cross_entropy.tex

spell:
	aspell -c cross_entropy.tex

s:
	make spell