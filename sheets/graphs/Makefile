graphs.tex: sheet.tex
	sheet -mode c -in sheet.tex > graphs.tex

graphs.pdf: ../../*.tex ../../trademark.pdf *.tex graphs.tex
	pdflatex graphs.tex
	make defs

make:
	make graphs.pdf

open: graphs.pdf
	open graphs.pdf

o:
	make open

defs: graphs.tex
	defs graphs.tex

spell:
	aspell -c graphs.tex

s:
	make spell