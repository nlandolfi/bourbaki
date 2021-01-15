inner_product_norms.tex: sheet.tex
	sheet -mode c -in sheet.tex > inner_product_norms.tex

inner_product_norms.pdf: ../../*.tex ../../trademark.pdf *.tex inner_product_norms.tex
	pdflatex inner_product_norms.tex
	make defs

make:
	make inner_product_norms.pdf

open: inner_product_norms.pdf
	open inner_product_norms.pdf

o:
	make open

defs: inner_product_norms.tex
	defs inner_product_norms.tex

spell:
	aspell -c inner_product_norms.tex

s:
	make spell