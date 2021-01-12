functional_analysis.tex: sheet.tex
	sheet -mode c -in sheet.tex > functional_analysis.tex

functional_analysis.pdf: ../../*.tex ../../trademark.pdf *.tex functional_analysis.tex
	pdflatex functional_analysis.tex
	make defs

make:
	make functional_analysis.pdf

open: functional_analysis.pdf
	open functional_analysis.pdf

o:
	make open

defs: functional_analysis.tex
	defs functional_analysis.tex

spell:
	aspell -c functional_analysis.tex

s:
	make spell