norm_examples.tex: sheet.tex
	sheet -mode c -in sheet.tex > norm_examples.tex

norm_examples.pdf: ../../*.tex ../../trademark.pdf *.tex norm_examples.tex
	pdflatex norm_examples.tex
	make defs

make:
	make norm_examples.pdf

open: norm_examples.pdf
	open norm_examples.pdf

o:
	make open

defs: norm_examples.tex
	defs norm_examples.tex

spell:
	aspell -c norm_examples.tex

s:
	make spell