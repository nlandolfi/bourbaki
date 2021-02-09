# generated automatically by sheet.go

make:
	make empirical_distribution.pdf

empirical_distribution.tex: sheet.tex
	sheet -mode c -in sheet.tex > empirical_distribution.tex

empirical_distribution.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_distribution.tex
	pdflatex empirical_distribution.tex
	make defs


open: empirical_distribution.pdf
	open empirical_distribution.pdf

o:
	make open

defs: empirical_distribution.tex
	defs empirical_distribution.tex

spell:
	aspell -c empirical_distribution.tex

s:
	make spell