metric_space_examples.tex: sheet.tex
	sheet -mode c -in sheet.tex > metric_space_examples.tex

metric_space_examples.pdf: ../../*.tex ../../trademark.pdf *.tex metric_space_examples.tex
	pdflatex metric_space_examples.tex
	make defs

make:
	make metric_space_examples.pdf

open: metric_space_examples.pdf
	open metric_space_examples.pdf

o:
	make open

defs: metric_space_examples.tex
	defs metric_space_examples.tex

spell:
	aspell -c metric_space_examples.tex

s:
	make spell