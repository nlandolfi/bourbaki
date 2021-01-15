complex_inner_products.tex: sheet.tex
	sheet -mode c -in sheet.tex > complex_inner_products.tex

complex_inner_products.pdf: ../../*.tex ../../trademark.pdf *.tex complex_inner_products.tex
	pdflatex complex_inner_products.tex
	make defs

make:
	make complex_inner_products.pdf

open: complex_inner_products.pdf
	open complex_inner_products.pdf

o:
	make open

defs: complex_inner_products.tex
	defs complex_inner_products.tex

spell:
	aspell -c complex_inner_products.tex

s:
	make spell