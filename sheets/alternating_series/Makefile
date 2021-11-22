# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make alternating_series.pdf

alternating_series.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > alternating_series.tex

alternating_series.pdf: ../../*.tex ../../trademark.pdf *.tex alternating_series.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode alternating_series.tex
	../../latexrun alternating_series.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: alternating_series.pdf
	open alternating_series.pdf

o:
	make open

terms: alternating_series.tex
	bbk_sheet -mode ts alternating_series.tex

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
	rm -f alternating_series.*
	make reset
	bbk_sheet -mode mk > Makefile