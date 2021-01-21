measure_space.tex: sheet.tex
	sheet -mode c -in sheet.tex > measure_space.tex

measure_space.pdf: ../../*.tex ../../trademark.pdf *.tex measure_space.tex
	pdflatex measure_space.tex
	make defs

make:
	make measure_space.pdf

open: measure_space.pdf
	open measure_space.pdf

o:
	make open

defs: measure_space.tex
	defs measure_space.tex

spell:
	aspell -c measure_space.tex

s:
	make spell