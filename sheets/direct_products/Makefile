# generated automatically by sheet.go

make:
	make direct_products.pdf

direct_products.tex: sheet.tex
	sheet -mode c -in sheet.tex > direct_products.tex

direct_products.pdf: ../../*.tex ../../trademark.pdf *.tex direct_products.tex
	pdflatex direct_products.tex
	make defs


open: direct_products.pdf
	open direct_products.pdf

o:
	make open

defs: direct_products.tex
	defs direct_products.tex

spell:
	aspell -c direct_products.tex

s:
	make spell