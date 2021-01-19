supremum_norm.tex: sheet.tex
	sheet -mode c -in sheet.tex > supremum_norm.tex

supremum_norm.pdf: ../../*.tex ../../trademark.pdf *.tex supremum_norm.tex
	pdflatex supremum_norm.tex
	make defs

make:
	make supremum_norm.pdf

open: supremum_norm.pdf
	open supremum_norm.pdf

o:
	make open

defs: supremum_norm.tex
	defs supremum_norm.tex

spell:
	aspell -c supremum_norm.tex

s:
	make spell