function_composition.tex: sheet.tex
	sheet -mode c -in sheet.tex > function_composition.tex

function_composition.pdf: ../../*.tex ../../trademark.pdf *.tex function_composition.tex
	pdflatex function_composition.tex
	make defs

make:
	make function_composition.pdf

open: function_composition.pdf
	open function_composition.pdf

o:
	make open

defs: function_composition.tex
	defs function_composition.tex

spell:
	aspell -c function_composition.tex

s:
	make spell