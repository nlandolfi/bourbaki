# generated automatically by sheet.go

make:
	make countable_probability_distributions.pdf

countable_probability_distributions.tex: sheet.tex
	sheet -mode c -in sheet.tex > countable_probability_distributions.tex

countable_probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex countable_probability_distributions.tex
	pdflatex countable_probability_distributions.tex
	make defs


open: countable_probability_distributions.pdf
	open countable_probability_distributions.pdf

o:
	make open

defs: countable_probability_distributions.tex
	defs countable_probability_distributions.tex

spell:
	aspell -c countable_probability_distributions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile