complete_metric_spaces.tex: sheet.tex
	sheet -mode c -in sheet.tex > complete_metric_spaces.tex

complete_metric_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex complete_metric_spaces.tex
	pdflatex complete_metric_spaces.tex
	make defs

make:
	make complete_metric_spaces.pdf

open: complete_metric_spaces.pdf
	open complete_metric_spaces.pdf

o:
	make open

defs: complete_metric_spaces.tex
	defs complete_metric_spaces.tex

spell:
	aspell -c complete_metric_spaces.tex

s:
	make spell