# generated automatically by sheet.go

make:
	make conditional_distributions.pdf

conditional_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > conditional_distributions.tex

conditional_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_distributions.tex
	pdflatex conditional_distributions.tex
	make defs


open: conditional_distributions.pdf
	open conditional_distributions.pdf

o:
	make open

defs: conditional_distributions.tex
	defs conditional_distributions.tex

spell:
	aspell -c conditional_distributions.tex

s:
	make spell