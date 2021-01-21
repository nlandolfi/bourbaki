metric_dense.tex: sheet.tex
	sheet -mode c -in sheet.tex > metric_dense.tex

metric_dense.pdf: ../../*.tex ../../trademark.pdf *.tex metric_dense.tex
	pdflatex metric_dense.tex
	make defs

make:
	make metric_dense.pdf

open: metric_dense.pdf
	open metric_dense.pdf

o:
	make open

defs: metric_dense.tex
	defs metric_dense.tex

spell:
	aspell -c metric_dense.tex

s:
	make spell