# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_completeness.pdf
	make graph.pdf

metric_completeness.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_completeness.tex

metric_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completeness.tex
	pdflatex --file-line-error -interaction=nonstopmode metric_completeness.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_completeness.pdf
	open metric_completeness.pdf

o:
	make open

terms: metric_completeness.tex
	bbk_sheet -mode ts metric_completeness.tex

spell:
	aspell -c metric_completeness.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f metric_completeness.*
	make reset
	bbk_sheet -mode mk > Makefile