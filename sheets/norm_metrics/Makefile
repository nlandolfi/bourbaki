norm_metrics.tex: sheet.tex
	sheet -mode c -in sheet.tex > norm_metrics.tex

norm_metrics.pdf: ../../*.tex ../../trademark.pdf *.tex norm_metrics.tex
	pdflatex norm_metrics.tex
	make defs

make:
	make norm_metrics.pdf

open: norm_metrics.pdf
	open norm_metrics.pdf

o:
	make open

defs: norm_metrics.tex
	defs norm_metrics.tex

spell:
	aspell -c norm_metrics.tex

s:
	make spell