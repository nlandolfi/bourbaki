# generated automatically by sheet.go

make:
	make distribution_expectation.pdf

distribution_expectation.tex: sheet.tex
	sheet -mode c -in sheet.tex > distribution_expectation.tex

distribution_expectation.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_expectation.tex
	pdflatex distribution_expectation.tex
	make defs


open: distribution_expectation.pdf
	open distribution_expectation.pdf

o:
	make open

defs: distribution_expectation.tex
	defs distribution_expectation.tex

spell:
	aspell -c distribution_expectation.tex

s:
	make spell