# generated automatically by sheet.go

make:
	make partial_orders.pdf

partial_orders.tex: sheet.tex
	sheet -mode c -in sheet.tex > partial_orders.tex

partial_orders.pdf: ../../*.tex ../../trademark.pdf *.tex partial_orders.tex
	pdflatex partial_orders.tex
	make defs


open: partial_orders.pdf
	open partial_orders.pdf

o:
	make open

defs: partial_orders.tex
	defs partial_orders.tex

spell:
	aspell -c partial_orders.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile