real_uniform_continuity.tex: sheet.tex
	sheet -mode c -in sheet.tex > real_uniform_continuity.tex

real_uniform_continuity.pdf: ../../*.tex ../../trademark.pdf *.tex real_uniform_continuity.tex
	pdflatex real_uniform_continuity.tex
	make defs

make:
	make real_uniform_continuity.pdf

open: real_uniform_continuity.pdf
	open real_uniform_continuity.pdf

o:
	make open

defs: real_uniform_continuity.tex
	defs real_uniform_continuity.tex

spell:
	aspell -c real_uniform_continuity.tex

s:
	make spell