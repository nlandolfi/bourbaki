# generated automatically by sheet.go

make:
	make differential_entropy.pdf

differential_entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > differential_entropy.tex

differential_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_entropy.tex
	pdflatex differential_entropy.tex
	make defs


open: differential_entropy.pdf
	open differential_entropy.pdf

o:
	make open

defs: differential_entropy.tex
	defs differential_entropy.tex

spell:
	aspell -c differential_entropy.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile