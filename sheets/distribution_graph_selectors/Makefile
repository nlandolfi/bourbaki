# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distribution_graph_selectors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

distribution_graph_selectors.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > distribution_graph_selectors.tex

distribution_graph_selectors.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_graph_selectors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distribution_graph_selectors.tex
	../../latexrun distribution_graph_selectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_graph_selectors.pdf
	open distribution_graph_selectors.pdf

o:
	make open

terms: distribution_graph_selectors.tex
	bbk_sheet -mode ts distribution_graph_selectors.tex

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
	rm -f distribution_graph_selectors.*
	make reset
	bbk_sheet -mode mk > Makefile