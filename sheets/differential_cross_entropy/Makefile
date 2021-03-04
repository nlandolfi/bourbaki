# generated automatically by sheet.go

make:
	make differential_cross_entropy.pdf

differential_cross_entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > differential_cross_entropy.tex

differential_cross_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_cross_entropy.tex
	pdflatex differential_cross_entropy.tex
	make defs


open: differential_cross_entropy.pdf
	open differential_cross_entropy.pdf

o:
	make open

defs: differential_cross_entropy.tex
	defs differential_cross_entropy.tex

spell:
	aspell -c differential_cross_entropy.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile