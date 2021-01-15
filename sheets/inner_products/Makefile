inner_products.tex: sheet.tex
	sheet -mode c -in sheet.tex > inner_products.tex

inner_products.pdf: ../../*.tex ../../trademark.pdf *.tex inner_products.tex
	pdflatex inner_products.tex
	make defs

make:
	make inner_products.pdf

open: inner_products.pdf
	open inner_products.pdf

o:
	make open

defs: inner_products.tex
	defs inner_products.tex

spell:
	aspell -c inner_products.tex

s:
	make spell