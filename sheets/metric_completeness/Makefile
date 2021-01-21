metric_completeness.tex: sheet.tex
	sheet -mode c -in sheet.tex > metric_completeness.tex

metric_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completeness.tex
	pdflatex metric_completeness.tex
	make defs

make:
	make metric_completeness.pdf

open: metric_completeness.pdf
	open metric_completeness.pdf

o:
	make open

defs: metric_completeness.tex
	defs metric_completeness.tex

spell:
	aspell -c metric_completeness.tex

s:
	make spell