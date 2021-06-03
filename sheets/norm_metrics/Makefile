# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make norm_metrics.pdf

norm_metrics.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > norm_metrics.tex

norm_metrics.pdf: ../../*.tex ../../trademark.pdf *.tex norm_metrics.tex
	# pdflatex --file-line-error -interaction=nonstopmode norm_metrics.tex
	../../latexrun norm_metrics.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_metrics.pdf
	open norm_metrics.pdf

o:
	make open

terms: norm_metrics.tex
	bbk_sheet -mode ts norm_metrics.tex

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
	rm -f norm_metrics.*
	make reset
	bbk_sheet -mode mk > Makefile