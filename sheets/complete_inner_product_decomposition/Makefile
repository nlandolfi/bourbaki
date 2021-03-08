# generated automatically by sheet.go

make:
	make complete_inner_product_decomposition.pdf

complete_inner_product_decomposition.tex: sheet.tex
	sheet -mode c -in sheet.tex > complete_inner_product_decomposition.tex

complete_inner_product_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex complete_inner_product_decomposition.tex
	pdflatex complete_inner_product_decomposition.tex
	make defs


open: complete_inner_product_decomposition.pdf
	open complete_inner_product_decomposition.pdf

o:
	make open

defs: complete_inner_product_decomposition.tex
	defs complete_inner_product_decomposition.tex

spell:
	aspell -c complete_inner_product_decomposition.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile