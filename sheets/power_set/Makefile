# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make power_set.pdf

power_set.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > power_set.tex

power_set.pdf: ../../*.tex ../../trademark.pdf *.tex power_set.tex
	# pdflatex --file-line-error -interaction=nonstopmode power_set.tex
	../../latexrun power_set.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: power_set.pdf
	open power_set.pdf

o:
	make open

terms: power_set.tex
	bbk_sheet -mode ts power_set.tex

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
	rm -f power_set.*
	make reset
	bbk_sheet -mode mk > Makefile