# generated automatically by sheet.go

make:
	make natural_numbers.pdf

natural_numbers.tex: sheet.tex
	sheet -mode c -in sheet.tex > natural_numbers.tex

natural_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex natural_numbers.tex
	pdflatex natural_numbers.tex
	make defs


open: natural_numbers.pdf
	open natural_numbers.pdf

o:
	make open

defs: natural_numbers.tex
	defs natural_numbers.tex

spell:
	aspell -c natural_numbers.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile