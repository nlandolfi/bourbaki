# generated automatically by sheet.go

make:
	make relative_entropy.pdf

relative_entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > relative_entropy.tex

relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex relative_entropy.tex
	pdflatex relative_entropy.tex
	make defs


open: relative_entropy.pdf
	open relative_entropy.pdf

o:
	make open

defs: relative_entropy.tex
	defs relative_entropy.tex

spell:
	aspell -c relative_entropy.tex

s:
	make spell