# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph_paths.pdf
	make graph.pdf

graph_paths.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > graph_paths.tex

graph_paths.pdf: ../../*.tex ../../trademark.pdf *.tex graph_paths.tex
	# pdflatex --file-line-error -interaction=nonstopmode graph_paths.tex
	../../latexrun graph_paths.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: graph_paths.pdf
	open graph_paths.pdf

o:
	make open

terms: graph_paths.tex
	bbk_sheet -mode ts graph_paths.tex

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
	rm -f graph_paths.*
	make reset
	bbk_sheet -mode mk > Makefile