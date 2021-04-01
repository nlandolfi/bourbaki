# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_completion.pdf
	make graph.pdf

metric_completion.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_completion.tex

metric_completion.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completion.tex
	pdflatex --file-line-error -interaction=nonstopmode metric_completion.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_completion.pdf
	open metric_completion.pdf

o:
	make open

terms: metric_completion.tex
	bbk_sheet -mode ts metric_completion.tex

spell:
	aspell -c metric_completion.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f metric_completion.*
	make reset
	bbk_sheet -mode mk > Makefile