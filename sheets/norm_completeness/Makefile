norm_completeness.tex: sheet.tex
	sheet -mode c -in sheet.tex > norm_completeness.tex

norm_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex norm_completeness.tex
	pdflatex norm_completeness.tex
	make defs

make:
	make norm_completeness.pdf

open: norm_completeness.pdf
	open norm_completeness.pdf

o:
	make open

defs: norm_completeness.tex
	defs norm_completeness.tex

spell:
	aspell -c norm_completeness.tex

s:
	make spell