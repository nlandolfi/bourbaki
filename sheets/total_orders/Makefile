# generated automatically by sheet.go

make:
	make total_orders.pdf

total_orders.tex: sheet.tex
	sheet -mode c -in sheet.tex > total_orders.tex

total_orders.pdf: ../../*.tex ../../trademark.pdf *.tex total_orders.tex
	pdflatex total_orders.tex
	make defs


open: total_orders.pdf
	open total_orders.pdf

o:
	make open

defs: total_orders.tex
	defs total_orders.tex

spell:
	aspell -c total_orders.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile