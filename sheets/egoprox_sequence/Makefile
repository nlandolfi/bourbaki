egoprox_sequence.tex: sheet.tex
	sheet -mode c -in sheet.tex > egoprox_sequence.tex

egoprox_sequence.pdf: ../../*.tex ../../trademark.pdf *.tex egoprox_sequence.tex
	pdflatex egoprox_sequence.tex
	make defs

make:
	make egoprox_sequence.pdf

open: egoprox_sequence.pdf
	open egoprox_sequence.pdf

o:
	make open

defs: egoprox_sequence.tex
	defs egoprox_sequence.tex

spell:
	aspell -c egoprox_sequence.tex

s:
	make spell