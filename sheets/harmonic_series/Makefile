# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make harmonic_series.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

harmonic_series.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > harmonic_series.tex

harmonic_series.pdf: ../../*.tex ../../trademark.pdf *.tex harmonic_series.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode harmonic_series.tex
	../../latexrun harmonic_series.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: harmonic_series.pdf
	open harmonic_series.pdf

o:
	make open

terms: harmonic_series.tex
	bbk sheet -mode ts harmonic_series.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f harmonic_series.*
	make reset
	bbk sheet -mode mk > Makefile