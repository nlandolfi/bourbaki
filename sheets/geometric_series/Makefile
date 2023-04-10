# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make geometric_series.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

geometric_series.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > geometric_series.tex

geometric_series.pdf: ../../*.tex ../../trademark.pdf *.tex geometric_series.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode geometric_series.tex
	../../latexrun geometric_series.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: geometric_series.pdf
	open geometric_series.pdf

o:
	make open

terms: geometric_series.tex
	bbk sheet -mode ts geometric_series.tex

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
	rm -f geometric_series.*
	make reset
	bbk sheet -mode mk > Makefile