absolute_value.tex: ../length/sheet.tex
	sheet -mode c -in ../length/sheet.tex > absolute_value.tex

absolute_value.pdf: ../../*.tex ../../trademark.pdf *.tex absolute_value.tex
	pdflatex absolute_value.tex
	make defs

make:
	make absolute_value.pdf

open: absolute_value.pdf
	open absolute_value.pdf

o:
	make open

defs: absolute_value.tex
	defs absolute_value.tex

spell:
	aspell -c absolute_value.tex

s:
	make spell