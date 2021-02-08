# generated automatically by sheet.go

make:
	make cross-entropy.pdf

cross-entropy.tex: sheet.tex
	sheet -mode c -in sheet.tex > cross-entropy.tex

cross-entropy.pdf: ../../*.tex ../../trademark.pdf *.tex cross-entropy.tex
	pdflatex cross-entropy.tex
	make defs


open: cross-entropy.pdf
	open cross-entropy.pdf

o:
	make open

defs: cross-entropy.tex
	defs cross-entropy.tex

spell:
	aspell -c cross-entropy.tex

s:
	make spell