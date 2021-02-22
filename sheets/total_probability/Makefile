# generated automatically by sheet.go

make:
	make total_probability.pdf

total_probability.tex: sheet.tex
	sheet -mode c -in sheet.tex > total_probability.tex

total_probability.pdf: ../../*.tex ../../trademark.pdf *.tex total_probability.tex
	pdflatex total_probability.tex
	make defs


open: total_probability.pdf
	open total_probability.pdf

o:
	make open

defs: total_probability.tex
	defs total_probability.tex

spell:
	aspell -c total_probability.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile