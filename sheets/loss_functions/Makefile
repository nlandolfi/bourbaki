# generated automatically by sheet.go

make:
	make loss_functions.pdf

loss_functions.tex: sheet.tex
	sheet -mode c -in sheet.tex > loss_functions.tex

loss_functions.pdf: ../../*.tex ../../trademark.pdf *.tex loss_functions.tex
	pdflatex loss_functions.tex
	make defs


open: loss_functions.pdf
	open loss_functions.pdf

o:
	make open

defs: loss_functions.tex
	defs loss_functions.tex

spell:
	aspell -c loss_functions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile