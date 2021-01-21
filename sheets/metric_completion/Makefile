metric_completion.tex: sheet.tex
	sheet -mode c -in sheet.tex > metric_completion.tex

metric_completion.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completion.tex
	pdflatex metric_completion.tex
	make defs

make:
	make metric_completion.pdf

open: metric_completion.pdf
	open metric_completion.pdf

o:
	make open

defs: metric_completion.tex
	defs metric_completion.tex

spell:
	aspell -c metric_completion.tex

s:
	make spell