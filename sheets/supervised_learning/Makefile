# generated automatically by sheet.go

make:
	make supervised_learning.pdf

supervised_learning.tex: sheet.tex
	sheet -mode c -in sheet.tex > supervised_learning.tex

supervised_learning.pdf: ../../*.tex ../../trademark.pdf *.tex supervised_learning.tex
	pdflatex supervised_learning.tex
	make defs


open: supervised_learning.pdf
	open supervised_learning.pdf

o:
	make open

defs: supervised_learning.tex
	defs supervised_learning.tex

spell:
	aspell -c supervised_learning.tex

s:
	make spell