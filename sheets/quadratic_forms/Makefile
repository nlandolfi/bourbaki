# generated automatically by sheet.go

make:
	make quadratic_forms.pdf

quadratic_forms.tex: sheet.tex
	sheet -mode c -in sheet.tex > quadratic_forms.tex

quadratic_forms.pdf: ../../*.tex ../../trademark.pdf *.tex quadratic_forms.tex
	pdflatex quadratic_forms.tex
	make defs


open: quadratic_forms.pdf
	open quadratic_forms.pdf

o:
	make open

defs: quadratic_forms.tex
	defs quadratic_forms.tex

spell:
	aspell -c quadratic_forms.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile