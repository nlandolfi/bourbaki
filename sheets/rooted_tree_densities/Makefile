# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rooted_tree_densities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

rooted_tree_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rooted_tree_densities.tex

rooted_tree_densities.pdf: ../../*.tex ../../trademark.pdf *.tex rooted_tree_densities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rooted_tree_densities.tex
	../../latexrun rooted_tree_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rooted_tree_densities.pdf
	open rooted_tree_densities.pdf

o:
	make open

terms: rooted_tree_densities.tex
	bbk_sheet -mode ts rooted_tree_densities.tex

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
	rm -f rooted_tree_densities.*
	make reset
	bbk_sheet -mode mk > Makefile