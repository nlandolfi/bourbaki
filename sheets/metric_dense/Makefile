# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_dense.pdf
	make graph.pdf

metric_dense.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_dense.tex

metric_dense.pdf: ../../*.tex ../../trademark.pdf *.tex metric_dense.tex
	# pdflatex --file-line-error -interaction=nonstopmode metric_dense.tex
	../../latexrun metric_dense.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_dense.pdf
	open metric_dense.pdf

o:
	make open

terms: metric_dense.tex
	bbk_sheet -mode ts metric_dense.tex

spell:
	aspell -c metric_dense.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f metric_dense.*
	make reset
	bbk_sheet -mode mk > Makefile