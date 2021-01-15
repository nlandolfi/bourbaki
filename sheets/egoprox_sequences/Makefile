egoprox_sequences.tex: sheet.tex
	sheet -mode c -in sheet.tex > egoprox_sequences.tex

egoprox_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex egoprox_sequences.tex
	pdflatex egoprox_sequences.tex
	make defs

make:
	make egoprox_sequences.pdf

open: egoprox_sequences.pdf
	open egoprox_sequences.pdf

o:
	make open

defs: egoprox_sequences.tex
	defs egoprox_sequences.tex

spell:
	aspell -c egoprox_sequences.tex

s:
	make spell