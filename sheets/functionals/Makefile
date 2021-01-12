functionals.tex: sheet.tex
	sheet -mode c -in sheet.tex > functionals.tex

functionals.pdf: ../../*.tex ../../trademark.pdf *.tex functionals.tex
	pdflatex functionals.tex
	make defs

make:
	make functionals.pdf

open: functionals.pdf
	open functionals.pdf

o:
	make open

defs: functionals.tex
	defs functionals.tex

spell:
	aspell -c functionals.tex

s:
	make spell