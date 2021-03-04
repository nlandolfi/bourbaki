# generated automatically by sheet.go

make:
	make probability_distributions.pdf

probability_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_distributions.tex

probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex probability_distributions.tex
	pdflatex probability_distributions.tex
	make defs


open: probability_distributions.pdf
	open probability_distributions.pdf

o:
	make open

defs: probability_distributions.tex
	defs probability_distributions.tex

spell:
	aspell -c probability_distributions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile