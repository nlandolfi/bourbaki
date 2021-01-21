function_inverses.tex: sheet.tex
	sheet -mode c -in sheet.tex > function_inverses.tex

function_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex function_inverses.tex
	pdflatex function_inverses.tex
	make defs

make:
	make function_inverses.pdf

open: function_inverses.pdf
	open function_inverses.pdf

o:
	make open

defs: function_inverses.tex
	defs function_inverses.tex

spell:
	aspell -c function_inverses.tex

s:
	make spell