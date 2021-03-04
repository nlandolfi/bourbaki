# generated automatically by sheet.go

make:
	make differential_relative_entropy.pdf

differential_relative_entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > differential_relative_entropy.tex

differential_relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex differential_relative_entropy.tex
	pdflatex differential_relative_entropy.tex
	make defs


open: differential_relative_entropy.pdf
	open differential_relative_entropy.pdf

o:
	make open

defs: differential_relative_entropy.tex
	defs differential_relative_entropy.tex

spell:
	aspell -c differential_relative_entropy.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile