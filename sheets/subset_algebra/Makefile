subset_algebra.tex: sheet.tex
	sheet -mode c -in sheet.tex > subset_algebra.tex

subset_algebra.pdf: ../../*.tex ../../trademark.pdf *.tex subset_algebra.tex
	pdflatex subset_algebra.tex
	make defs

make:
	make subset_algebra.pdf

open: subset_algebra.pdf
	open subset_algebra.pdf

o:
	make open

defs: subset_algebra.tex
	defs subset_algebra.tex

spell:
	aspell -c subset_algebra.tex

s:
	make spell