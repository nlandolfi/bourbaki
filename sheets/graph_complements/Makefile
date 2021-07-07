# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make graph_complements.pdf

graph_complements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > graph_complements.tex

graph_complements.pdf: ../../*.tex ../../trademark.pdf *.tex graph_complements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode graph_complements.tex
	../../latexrun graph_complements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: graph_complements.pdf
	open graph_complements.pdf

o:
	make open

terms: graph_complements.tex
	bbk_sheet -mode ts graph_complements.tex

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
	rm -f graph_complements.*
	make reset
	bbk_sheet -mode mk > Makefile