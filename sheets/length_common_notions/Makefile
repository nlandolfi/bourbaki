# generated automatically by sheet.go

make:
	make length_common_notions.pdf

length_common_notions.tex: sheet.tex
	sheet -mode c -in sheet.tex > length_common_notions.tex

length_common_notions.pdf: ../../*.tex ../../trademark.pdf *.tex length_common_notions.tex
	pdflatex length_common_notions.tex
	make defs


open: length_common_notions.pdf
	open length_common_notions.pdf

o:
	make open

defs: length_common_notions.tex
	defs length_common_notions.tex

spell:
	aspell -c length_common_notions.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile