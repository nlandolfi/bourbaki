# generated automatically by sheet.go

make:
	make n-dimensional_space.pdf

n-dimensional_space.tex: sheet.tex
	sheet -mode c -in sheet.tex > n-dimensional_space.tex

n-dimensional_space.pdf: ../../*.tex ../../trademark.pdf *.tex n-dimensional_space.tex
	pdflatex n-dimensional_space.tex
	make defs


open: n-dimensional_space.pdf
	open n-dimensional_space.pdf

o:
	make open

defs: n-dimensional_space.tex
	defs n-dimensional_space.tex

spell:
	aspell -c n-dimensional_space.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile