# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directed_graph_skeletons.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

directed_graph_skeletons.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > directed_graph_skeletons.tex

directed_graph_skeletons.pdf: ../../*.tex ../../trademark.pdf *.tex directed_graph_skeletons.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directed_graph_skeletons.tex
	../../latexrun directed_graph_skeletons.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directed_graph_skeletons.pdf
	open directed_graph_skeletons.pdf

o:
	make open

terms: directed_graph_skeletons.tex
	bbk_sheet -mode ts directed_graph_skeletons.tex

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
	rm -f directed_graph_skeletons.*
	make reset
	bbk_sheet -mode mk > Makefile